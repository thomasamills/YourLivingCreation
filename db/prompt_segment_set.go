package db

import (
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) GetPromptSegmentSet(
	promptSegmentSetId string,
	upperTx *gorm.DB,
) ([]*humanize_protobuf.PromptSegment, error) {
	promptSegmentSet := make([]*humanize_protobuf.PromptSegment, 0)
	getPromptSegmentSet := func(tx *gorm.DB) error {
		promptSegmentIds, err := h.ListPromptSegments(promptSegmentSetId, tx)
		if err != nil {
			return err
		}
		for _, id := range promptSegmentIds {
			promptSegment, err := h.GetPromptSegment(id, tx)
			if err != nil {
				return err
			}
			promptSegmentSet = append(promptSegmentSet, promptSegment)
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = getPromptSegmentSet(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = getPromptSegmentSet(tx)
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
	return promptSegmentSet, nil
}

func (h *HumanizeDbImpl) ListPromptSegmentSets(
	upperTx *gorm.DB,
) ([]string, error) {
	var arr []string
	promptSegmentSetIds := func(tx *gorm.DB) error {
		err := tx.Model(PromptSegment{}).
			Select("prompt_segment__set_id").
			Distinct().
			Find(&arr).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = promptSegmentSetIds(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = promptSegmentSetIds(tx)
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
