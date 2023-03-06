package db

import (
	"fmt"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) ListActuationRules(
	personalityId string,
	upperTx *gorm.DB,
) ([]string, error) {
	var arr []string
	listActuationRules := func(tx *gorm.DB) error {
		err := tx.Model(ActuationRule{}).Where(
			"personality_id = ?", personalityId,
		).Select("actuation_rule_id").Find(&arr).Error
		if err != nil {
			return err
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = listActuationRules(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = listActuationRules(tx)
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

func (h *HumanizeDbImpl) CreateActuationRuleSet(
	personalityId string,
	actuationRuleSet []*humanize_protobuf.ActuationRule,
) error {
	err := h.mainDB.Transaction(func(tx *gorm.DB) error {
		for _, actuationRule := range actuationRuleSet {
			err := h.CreateOrUpdateActuationRule(
				actuationRule.ActuationRuleId, actuationRule, tx, personalityId,
			)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *HumanizeDbImpl) CreateOrUpdateActuationRule(
	actuationRulePartId string,
	rule *humanize_protobuf.ActuationRule,
	upperLevelTx *gorm.DB,
	personalityId string,
) error {
	err := h.mainDB.Transaction(func(tx *gorm.DB) error {
		if upperLevelTx != nil {
			tx = upperLevelTx
		}
		actuationRule := &ActuationRule{
			RuleType:                 int32(rule.Type),
			ActuationRuleId:          rule.ActuationRuleId,
			PersonalityId:            personalityId,
			ResultingPromptId:        rule.ResultingPromptId,
			ResultingPromptSegmentId: rule.ResultingPromptSegmentId,
			IdealStateId:             rule.IdealEmotionalState.EntityId,
			RuleName:                 rule.RuleName,
		}
		txErr := tx.Where("actuation_rule_part_id = ?", actuationRulePartId).
			FirstOrCreate(actuationRule).
			Updates(actuationRule)

		for _, rulePart := range rule.RuleParts {
			// assign rule part ids
			_, err := h.CreateEmotionalBoundRulePart(rulePart)
			if err != nil {
				return err
			}
		}
		if txErr != nil {
			tx.Rollback()
			return txErr.Error
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *HumanizeDbImpl) GetActuationRuleSet(
	personalityId string,
	upperTx *gorm.DB,
) ([]*humanize_protobuf.ActuationRule, error) {
	actuationRules := make([]*humanize_protobuf.ActuationRule, 0)

	getActuationRules := func(tx *gorm.DB) error {
		rule, err := h.GetActuationRule(personalityId, tx, true)
		if err != nil {
			return err
		}
		actuationRules = append(actuationRules, rule)
		return nil
	}
	var err error
	if upperTx != nil {
		err = getActuationRules(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			return getActuationRules(upperTx)
		})
	}
	if err != nil {
		return nil, err
	}
	return actuationRules, nil
}

func (h *HumanizeDbImpl) GetActuationRule(
	id string, upperTx *gorm.DB,
	byPersonalityId bool,
) (*humanize_protobuf.ActuationRule, error) {
	actuationRule := &humanize_protobuf.ActuationRule{}
	err := h.mainDB.Transaction(func(tx *gorm.DB) error {
		// First obtain the single rule
		if upperTx != nil {
			tx = upperTx
		}
		var rule = ActuationRule{}
		if byPersonalityId {
			err := tx.
				Where("personality_id = ?", id).
				First(&rule).
				Error
			if err != nil {
				return err
			}
		} else {
			err := tx.
				Where("actuation_rule_id = ?", id).
				First(&rule).
				Error
			if err != nil {
				return err
			}
		}

		actuationRule.ActuationRuleId = rule.ActuationRuleId
		actuationRule.PercentageOfProc = rule.PercentageOfProc
		actuationRule.Type = humanize_protobuf.ActuationRuleType(rule.RuleType)
		// Ideal emotional state retrieveal
		state, err := h.GetEmotionalState(rule.IdealStateId, nil)
		if err != nil {
			return err
		}
		actuationRule.IdealEmotionalState = state
		actuationRule.PersonalityId = rule.PersonalityId

		// Now get the rule parts
		var ruleParts []RulePart
		err = tx.
			Where("personality_id = ?", rule.PersonalityId).
			Find(&ruleParts).
			Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println(err.Error())
			}
			return err
		}
		for _, rulePart := range ruleParts {
			actuationRule.RuleParts = append(actuationRule.RuleParts, &humanize_protobuf.RulePart{
				PercentageOfProc: rulePart.PercentageOfProc,
				RuleId:           rulePart.RuleID,
				RulePartId:       rulePart.RulePartId,
				Type:             humanize_protobuf.RulePart_RulePartType(rulePart.RuleType),
				BoundPercentage:  rulePart.BoundPercentage,
				BoundId:          rulePart.BoundId,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return actuationRule, nil

}

func (h *HumanizeDbImpl) DeleteActuationRule(actuationRuleId string) error {
	err := h.mainDB.Where("actuation_rule_id = ?", actuationRuleId).Delete(&ActuationRule{})
	if err != nil {
		if err.Error != nil {
			return err.Error
		}
	}
	return nil
}
