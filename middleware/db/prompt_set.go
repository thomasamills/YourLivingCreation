package db

import (
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) GetPromptSet(
	promptSetId string,
	upperTx *gorm.DB,
) ([]*humanize_protobuf.Prompt, error) {
	prompts := make([]*humanize_protobuf.Prompt, 0)
	getPromptSet := func(tx *gorm.DB) error {
		promptIds, err := h.ListPrompts(promptSetId, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, promptId := range promptIds {
			prompt, err := h.GetPrompt(promptId, tx)
			if err != nil {
				return err
			}
			prompts = append(prompts, prompt)
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = getPromptSet(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			return getPromptSet(tx)
		})
	}
	if err != nil {
		return nil, err
	}
	return prompts, nil
}

func (h *HumanizeDbImpl) GetPromptIdealStatesBySetId(
	promptSetId string,
	upperTx *gorm.DB,
) ([]IdealStateAndId, error) {
	idAndStates := make([]IdealStateAndId, 0)
	getPromptSet := func(tx *gorm.DB) error {
		promptIds, err := h.ListPrompts(promptSetId, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, promptId := range promptIds {
			emotionalState, err := h.GetEmotionalState(promptId, tx)
			if err != nil {
				return err
			}
			idAndStates = append(idAndStates, IdealStateAndId{
				Id:    promptId,
				State: emotionalState,
			})
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = getPromptSet(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			return getPromptSet(tx)
		})
	}
	if err != nil {
		return nil, err
	}
	return idAndStates, nil
}

func (h *HumanizeDbImpl) ListPromptSetIds(
	upperTx *gorm.DB,
) ([]string, error) {
	var arr []string
	listPrompts := func(tx *gorm.DB) error {
		err := tx.Model(Prompt{}).Select("prompt_set_id").Distinct().Find(&arr).Error
		if err != nil {
			return err
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = listPrompts(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = listPrompts(tx)
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
