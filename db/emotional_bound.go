package db

import (
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreateEmotionalBound(
	bound *humanize_protobuf.EmotionalBound,
	entityId string, // if we are overriding the entity id
	upperTx *gorm.DB,
) (string, error) {
	boundInstanceId := h.idGen.GetULIDfromtime()
	if len(entityId) == 0 {
		entityId = bound.EntityId
	}
	if len(entityId) == 0 {
		entityId = bound.EntityId
	}
	createEmotionalBound := func(tx *gorm.DB) error {
		b := &EmotionalBound{
			BoundInstanceId:         boundInstanceId,
			EntityId:                entityId,
			BoundId:                 bound.BoundId,
			CurrentPercentage:       int64(bound.CurrentPercentage),
			LowerBound:              bound.LowerBound,
			UpperBound:              bound.UpperBound,
			EncodedSynonymMap:       h.encodeSynonymMap(bound.Synonyms),
			EmotionalShiftMagnitude: int32(bound.EmotionShiftMagnitude),
		}
		if err := tx.Create(b); err.Error != nil {
			return err.Error
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = createEmotionalBound(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = createEmotionalBound(tx)
			if err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
		if err != nil {
			return "", err
		}

	}

	return boundInstanceId, nil
}

func (h *HumanizeDbImpl) UpdateEmotionalBound(
	boundInstanceId string, bound *humanize_protobuf.EmotionalBound,
) (string, error) {
	b := &EmotionalBound{
		BoundInstanceId:         boundInstanceId,
		EntityId:                bound.EntityId,
		BoundId:                 bound.BoundId,
		CurrentPercentage:       int64(bound.CurrentPercentage),
		LowerBound:              bound.LowerBound,
		UpperBound:              bound.UpperBound,
		EncodedSynonymMap:       h.encodeSynonymMap(bound.Synonyms),
		EmotionalShiftMagnitude: int32(bound.EmotionShiftMagnitude),
	}
	err := h.mainDB.Where("bound_instance_id = ?", boundInstanceId).Updates(b)
	if err.Error != nil {
		return "", err.Error
	}
	return b.BoundInstanceId, nil
}

func (h *HumanizeDbImpl) UpdateEmotionalBoundPercentage(
	emotionalBound *humanize_protobuf.EmotionalBound,
) error {
	eb := &EmotionalBound{
		BoundInstanceId: emotionalBound.BoundInstanceId,
	}
	err := h.mainDB.Model(&eb).Updates(EmotionalBound{
		CurrentPercentage: int64(emotionalBound.CurrentPercentage),
	})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (h *HumanizeDbImpl) DeleteEmotionalBound(
	boundInstanceId string,
	upperTx *gorm.DB,
) error {
	deleteBound := func(tx *gorm.DB) error {
		err := h.mainDB.Where("bound_instance_id = ?", boundInstanceId).Delete(&EmotionalBound{})
		if err != nil {
			if err.Error != nil {
				return err.Error
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = deleteBound(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = deleteBound(tx)
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
