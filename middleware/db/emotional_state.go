package db

import (
	"errors"
	"fmt"
	"strings"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreateEmotionalState(
	entityId string,
	isPreset bool,
	state *humanize_protobuf.EmotionalState,
	upperTx *gorm.DB,
) error {
	saveEmotionalState := func(tx *gorm.DB) error {
		emotionalState := &EmotionalState{
			IsPreset:      isPreset,
			EntityId:      entityId,
			PersonalityID: state.PersonalityId,
			Composure:     state.Composure,
			Likeness:      state.Likeness,
			Fears:         strings.Join(state.Fears, ","),
			Hobbies:       strings.Join(state.Hobbies, ","),
		}
		txErr := tx.Where("entity_id = ?", entityId).
			Create(emotionalState)
		if txErr.Error != nil {
			return txErr.Error
		}
		for _, boundary := range state.EmotionalBounds {
			_, err := h.CreateEmotionalBound(boundary, entityId, tx)
			if err != nil {
				return err
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = saveEmotionalState(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = saveEmotionalState(tx)
			if err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	}
	if err != nil {
		return err
	}
	return nil
}

func (h *HumanizeDbImpl) CopyEmotionalState(
	presetId, entityId string,
) (*humanize_protobuf.EmotionalState, error) {
	presetState, err := h.GetEmotionalState(presetId, nil)
	if err != nil {
		return nil, err
	}
	if !presetState.IsPreset {
		return nil, errors.New("cannot copy emotional state, state is not preset")
	}
	if err := h.CreateEmotionalState(entityId, false, presetState, nil); err != nil {
		return nil, err
	}
	// TODO better memory handling here
	return h.GetEmotionalState(entityId, nil)
}

func (h *HumanizeDbImpl) GetEmotionalState(
	entityId string,
	upperTx *gorm.DB,
) (*humanize_protobuf.EmotionalState, error) {
	protoState := &humanize_protobuf.EmotionalState{}
	getEmotionalState := func(tx *gorm.DB) error {
		emotionalState := &EmotionalState{}
		err := tx.
			Where("entity_id = ?", entityId).
			First(&emotionalState).
			Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println(err.Error())
			}
			return err
		}
		var bounds []EmotionalBound
		err = tx.
			Where("entity_id = ?", entityId).
			Find(&bounds).
			Error
		if err != nil {
			return err
		}
		protoState = &humanize_protobuf.EmotionalState{
			IsPreset:        emotionalState.IsPreset,
			EntityId:        emotionalState.EntityId,
			PersonalityId:   emotionalState.PersonalityID,
			Composure:       emotionalState.Composure,
			Likeness:        emotionalState.Likeness,
			Fears:           strings.Split(emotionalState.Fears, ","),
			Hobbies:         strings.Split(emotionalState.Hobbies, ","),
			EmotionalBounds: make(map[string]*humanize_protobuf.EmotionalBound, 0),
		}
		for _, bound := range bounds {
			decodedSynonymMap, err := h.decodeSynonymMap(bound.EncodedSynonymMap)
			if err != nil {
				return err
			}
			protoState.EmotionalBounds[bound.BoundId] = &humanize_protobuf.EmotionalBound{
				EntityId:              emotionalState.EntityId,
				BoundId:               bound.BoundId,
				BoundInstanceId:       bound.BoundInstanceId,
				CurrentPercentage:     int32(bound.CurrentPercentage),
				LowerBound:            bound.LowerBound,
				UpperBound:            bound.UpperBound,
				Synonyms:              decodedSynonymMap,
				EmotionShiftMagnitude: humanize_protobuf.EmotionShiftMagnitude(bound.EmotionalShiftMagnitude),
			}
		}
		return nil
	}

	var err error
	if upperTx != nil {
		err = getEmotionalState(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = getEmotionalState(tx)
			if err != nil {
				tx.Rollback()
				return err
			}

			return nil
		})
	}
	if err != nil {
		return nil, err
	}
	return protoState, nil
}
func (h *HumanizeDbImpl) DeleteEmotionalState(
	entityId string,
	upperTx *gorm.DB,
) error {
	deleteEmotionalState := func(tx *gorm.DB) error {
		err := tx.Where("entity_id = ?", entityId).Delete(&EmotionalState{})
		if err != nil {
			if err.Error != nil {
				tx.Rollback()
				return err.Error
			}
		}
		// Then list emotional bounds for state and delete
		return nil
	}
	var err error
	if upperTx != nil {
		err = deleteEmotionalState(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = deleteEmotionalState(tx)
			if err != nil {
				tx.Rollback()
				return err
			}

			return nil

		})
	}
	if err != nil {
		return err
	}
	return nil
}
