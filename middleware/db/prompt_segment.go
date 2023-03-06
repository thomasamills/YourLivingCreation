package db

import (
	"fmt"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreatePromptSegment(
	promptSegment *humanize_protobuf.PromptSegment,
	upperTx *gorm.DB,
) (string, error) {
	promptSegmentId := promptSegment.PromptSegmentId
	createPromptSegment := func(tx *gorm.DB) error {
		if len(promptSegmentId) == 0 {
			promptSegmentId = h.idGen.GetULIDfromtime()
		}
		p := &PromptSegment{
			Message:            promptSegment.Message,
			PromptSegmentId:    promptSegmentId,
			PromptSegmentSetId: promptSegment.PromptSegmentSetId,
			PromptSegmentType:  int32(promptSegment.Type),
			PercentageOfProc:   promptSegment.PercentageOfProc,
			PrimerType:         promptSegment.PrimerType,
		}
		err := tx.Create(p)
		if err != nil {
			if err.Error != nil {
				return err.Error
			}
		}
		rulePartErr := h.CreateRuleParts(p.PromptSegmentId, promptSegment.RuleParts, tx)
		if rulePartErr != nil {
			return rulePartErr
		}

		idealEmotionalStateErr := h.CreateEmotionalState(
			p.PromptSegmentId, true, promptSegment.IdealEmotionalState, tx,
		)

		if idealEmotionalStateErr != nil {
			return idealEmotionalStateErr
		}
		return nil
	}
	var txExecuteErr error
	if upperTx != nil {
		txExecuteErr = createPromptSegment(upperTx)
	} else {
		txExecuteErr = h.mainDB.Transaction(func(tx *gorm.DB) error {
			txExecuteErr = createPromptSegment(tx)
			if txExecuteErr != nil {
				tx.Rollback()
				return txExecuteErr
			}

			return nil
		})
	}
	if txExecuteErr != nil {
		return "", txExecuteErr
	}
	return promptSegmentId, nil
}

func (h *HumanizeDbImpl) ListPromptSegmentsBySetId(
	promptSegmentSetId string,
	upperTx *gorm.DB,
) ([]string, error) {
	var arr []string
	listPromptSegments := func(tx *gorm.DB) error {
		err := tx.Model(PromptSegment{}).
			Where("prompt_segment_set_id = ?", promptSegmentSetId).
			Select("prompt_segment_id").
			Find(&arr).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = listPromptSegments(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = listPromptSegments(tx)
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
	return arr, nil
}

func (h *HumanizeDbImpl) GetPromptSegment(
	promptSegmentId string,
	upperTx *gorm.DB,
) (*humanize_protobuf.PromptSegment, error) {
	segment := &humanize_protobuf.PromptSegment{}
	getPromptSegment := func(tx *gorm.DB) error {
		result := &PromptSegment{}
		err := h.mainDB.
			Where("prompt_segment_id = ?", promptSegmentId).
			First(&result).
			Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println(err.Error())
			}
			return err
		}
		segment.Type = humanize_protobuf.PromptSegmentType(result.PromptSegmentType)
		segment.PromptSegmentId = result.PromptSegmentId
		segment.PromptSegmentSetId = result.PromptSegmentSetId
		segment.Message = result.Message
		segment.PercentageOfProc = result.PercentageOfProc
		segment.PrimerType = result.PrimerType
		// Get the rules
		ruleParts, err := h.GetRuleParts(segment.PromptSegmentId, tx)
		if err != nil {
			return err
		}
		segment.RuleParts = ruleParts
		// now get emotional state
		idealEmotionalState, err := h.GetEmotionalState(promptSegmentId, tx)
		if err != nil {
			return err
		}
		segment.IdealEmotionalState = idealEmotionalState
		return nil
	}
	var err error
	if upperTx != nil {
		err = getPromptSegment(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = getPromptSegment(tx)
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
	return segment, nil
}

func (h *HumanizeDbImpl) UpdatePromptSegment(
	promptSegment *humanize_protobuf.PromptSegment,
	upperTx *gorm.DB,
) error {
	err := h.mainDB.Transaction(func(tx *gorm.DB) error {
		p := &PromptSegment{
			Message:            promptSegment.Message,
			PromptSegmentId:    promptSegment.PromptSegmentId,
			PromptSegmentSetId: promptSegment.PromptSegmentSetId,
			PromptSegmentType:  int32(promptSegment.Type),
			PrimerType:         promptSegment.PrimerType,
		}
		err := h.mainDB.Updates(p)
		if err != nil {
			if err.Error != nil {
				return err.Error
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *HumanizeDbImpl) DeletePromptSegment(
	promptId string,
	upperTx *gorm.DB,
) error {
	deletePromptSegment := func(tx *gorm.DB) error {
		err := tx.Where("prompt_segment_id = ?", promptId).Delete(&PromptSegment{})
		if err != nil {
			if err.Error != nil {
				return err.Error
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = deletePromptSegment(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = deletePromptSegment(tx)
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
