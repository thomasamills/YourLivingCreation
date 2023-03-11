package db

import (
	"errors"
	"fmt"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreatePrompt(
	prompt *humanize_protobuf.Prompt,
	upperTx *gorm.DB,
) (string, error) {
	if prompt == nil {
		return "", errors.New("request contains no prompt")
	}
	success, errMessage := h.validatePromptRequest(prompt)
	if !success {
		return "", errors.New(errMessage)
	}
	promptId := prompt.PromptId

	createPrompt := func(tx *gorm.DB) error {
		if len(promptId) == 0 {
			promptId = h.idGen.GetULIDfromtime()
		}
		// Get the rules
		// create rule parts now.
		err := h.CreateRuleParts(promptId, prompt.RuleParts, tx)
		if err != nil {
			return err
		}

		p := &Prompt{
			PromptId:             promptId,
			PromptText:           prompt.PromptText,
			PromptName:           prompt.PromptName,
			PromptType:           int32(prompt.PromptType),
			PromptSetId:          prompt.PromptSetId,
			RequiredSegmentTypes: h.encodePromptSegmentTypes(prompt.RequiredPromptSegmentTypes),
			StoryBackground:      prompt.StoryBackground,
		}

		idealEmotionalStateErr := h.CreateEmotionalState(
			p.PromptId, true, prompt.IdealEmotionalState, tx,
		)

		if idealEmotionalStateErr != nil {
			return idealEmotionalStateErr
		}

		txErr := tx.Create(p)
		if txErr != nil {
			if txErr.Error != nil {
				if txErr.Error != gorm.ErrRecordNotFound {
					return txErr.Error
				}
			}
		}
		return nil
	}

	var err error
	if upperTx != nil {
		err = createPrompt(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = createPrompt(tx)
			if err != nil {
				tx.Rollback()
				return err
			}

			return nil
		})
	}
	if err != nil {
		return "", err
	}
	return promptId, nil
}

func (h *HumanizeDbImpl) GetPrompt(
	promptId string,
	upperTx *gorm.DB,
) (*humanize_protobuf.Prompt, error) {
	if len(promptId) == 0 {
		return nil, errors.New("failed to load the prompt from the database, prompt_id was nil")
	}
	prompt := &humanize_protobuf.Prompt{}
	getPrompt := func(tx *gorm.DB) error {
		result := &Prompt{}
		err := tx.
			Where("prompt_id = ?", promptId).
			First(&result).
			Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println(err.Error())
			}
			return err
		}
		prompt.PromptType = humanize_protobuf.PromptType(result.PromptType)
		prompt.PromptId = result.PromptId
		prompt.PromptName = result.PromptName
		prompt.PromptSetId = result.PromptSetId
		prompt.PromptText = result.PromptText
		prompt.StoryBackground = result.StoryBackground

		requiredPromptSegments, err := h.decodePromptSegmentTypes(result.RequiredSegmentTypes)
		if err != nil {
			return err
		}
		prompt.RequiredPromptSegmentTypes = requiredPromptSegments
		// Get the rules
		ruleParts, err := h.GetRuleParts(promptId, tx)
		if err != nil {
			return err
		}
		for _, rulePart := range ruleParts {
			prompt.RuleParts = append(prompt.RuleParts, rulePart)
		}

		// now get the ideal emotional state
		emotionalState, err := h.GetEmotionalState(promptId, tx)
		if err != nil {
			return nil
		}
		prompt.IdealEmotionalState = emotionalState
		return nil
	}
	var err error
	if upperTx != nil {
		err = getPrompt(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = getPrompt(tx)
			if err != nil {
				tx.Rollback()
				return nil
			}

			return nil
		})
	}

	if err != nil {
		return nil, err
	}

	return prompt, nil
}

func (h *HumanizeDbImpl) ListPrompts(
	promptSetId string,
	upperTx *gorm.DB,
) ([]string, error) {
	var arr []string
	listPrompts := func(tx *gorm.DB) error {
		err := tx.Model(Prompt{}).Where(
			"prompt_set_id = ?", promptSetId,
		).Select("prompt_id").Find(&arr).Error
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

func (h *HumanizeDbImpl) UpdatePrompt(
	prompt *humanize_protobuf.Prompt,
	upperTx *gorm.DB,
) error {
	updatePrompt := func(tx *gorm.DB) error {
		err := tx.Updates(prompt)
		if err != nil {
			if err.Error != nil {
				return err.Error
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = updatePrompt(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = updatePrompt(tx)
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

func (h *HumanizeDbImpl) DeletePrompt(
	promptId string,
	upperTx *gorm.DB,
) error {
	deletePrompt := func(tx *gorm.DB) error {
		err := tx.Where("prompt_id = ?", promptId).Delete(&Prompt{})
		if err != nil {
			if err.Error != nil {
				return err.Error
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = deletePrompt(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = deletePrompt(tx)
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
