package db

import (
	"fmt"
	"strings"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreateRuleParts(
	associatingId string,
	ruleParts []*humanize_protobuf.RulePart,
	upperTx *gorm.DB,
) error {
	createRuleParts := func(tx *gorm.DB) error {
		for _, protoRulePart := range ruleParts {
			rulePart := &RulePart{
				RulePartId:               h.idGen.GetULIDfromtime(),
				RuleID:                   associatingId,
				PercentageOfProc:         protoRulePart.PercentageOfProc,
				RuleType:                 int32(protoRulePart.Type),
				BoundPercentage:          protoRulePart.BoundPercentage,
				BoundId:                  protoRulePart.BoundId,
				ConversationEmotionTrain: strings.Join(protoRulePart.ConversationEmotionTrain, ","),
			}
			if txErr := tx.Where("rule_id = ?", associatingId).
				Create(rulePart); txErr.Error != nil {
				return txErr.Error
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = createRuleParts(upperTx)
		return nil
	}
	err = h.mainDB.Transaction(func(tx *gorm.DB) error {
		err = createRuleParts(tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *HumanizeDbImpl) GetRulePart(
	rulePartId string, upperTx *gorm.DB,
) (*humanize_protobuf.RulePart, error) {
	protoRulePart := &humanize_protobuf.RulePart{}
	getRulePart := func(tx *gorm.DB) error {
		var rulePart RulePart
		err := tx.
			Where("rule_part_id = ?", rulePartId).
			Find(&rulePart).
			Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println(err.Error())
			}
			return err
		}

		protoRulePart = &humanize_protobuf.RulePart{
			PercentageOfProc:         rulePart.PercentageOfProc,
			BoundPercentage:          rulePart.BoundPercentage,
			BoundId:                  rulePart.BoundId,
			Type:                     humanize_protobuf.RulePart_RulePartType(rulePart.RuleType),
			RulePartId:               rulePart.RulePartId,
			ConversationEmotionTrain: strings.Split(rulePart.ConversationEmotionTrain, ","),
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = getRulePart(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			return getRulePart(tx)
		})
	}
	if err != nil {
		return nil, err
	}
	return protoRulePart, nil
}

func (h *HumanizeDbImpl) GetRuleParts(
	ruleId string,
	upperTx *gorm.DB,
) ([]*humanize_protobuf.RulePart, error) {
	result := make([]*humanize_protobuf.RulePart, 0)
	err := h.mainDB.Transaction(func(tx *gorm.DB) error {
		if upperTx != nil {
			tx = upperTx
		}
		var ruleParts []RulePart
		err := tx.
			Where("rule_id = ?", ruleId).
			Find(&ruleParts).
			Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println(err.Error())
			}
			return err
		}

		for _, rulePart := range ruleParts {
			result = append(result, &humanize_protobuf.RulePart{
				PercentageOfProc:         rulePart.PercentageOfProc,
				BoundPercentage:          rulePart.BoundPercentage,
				BoundId:                  rulePart.BoundId,
				Type:                     humanize_protobuf.RulePart_RulePartType(rulePart.RuleType),
				RulePartId:               rulePart.RulePartId,
				ConversationEmotionTrain: strings.Split(rulePart.ConversationEmotionTrain, ","),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (h *HumanizeDbImpl) DeletePersonalityRule(
	ruleId string,
	upperTx *gorm.DB,
) error {
	deletePersonalityRule := func(tx *gorm.DB) error {
		err := h.mainDB.Where("rule_id = ?", ruleId).Delete(&EmotionalBoundRule{})
		if err != nil {
			if err.Error != nil {
				return err.Error
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = deletePersonalityRule(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			return deletePersonalityRule(upperTx)
		})
	}
	if err != nil {
		return err
	}
	return nil
}

func (h *HumanizeDbImpl) DeletePersonalityRulePart(
	ruleId string,
	upperTx *gorm.DB,
) error {
	deletePersonalityRule := func(tx *gorm.DB) error {
		err := h.mainDB.Where("rule_part_id = ?", ruleId).Delete(&RulePart{})
		if err != nil {
			if err.Error != nil {
				return err.Error
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = deletePersonalityRule(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			return deletePersonalityRule(upperTx)
		})
	}
	if err != nil {
		return err
	}
	return nil
}
