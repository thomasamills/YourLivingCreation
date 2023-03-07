package db

import (
	"fmt"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) SavePersonality(
	personalityId string,
	rules []humanize_protobuf.EmotionUpdateRule,
	upperTx *gorm.DB,
) error {
	savePersonality := func(tx *gorm.DB) error {
		personality := &Personality{
			PersonalityId: personalityId,
		}
		txErr := tx.Create(personality)
		if txErr != nil {
			if tx.Error != nil {
				return nil
			}
		}
		for _, protoRule := range rules {
			// generate personality id
			id := h.idGen.GetULIDfromtime()
			// encode bound shifts
			rule := &EmotionalBoundRule{
				RuleID:                       id,
				PersonalityID:                personalityId,
				Name:                         protoRule.Name,
				PercentageOfProc:             protoRule.PercentageOfProc,
				TriggeringAction:             protoRule.TriggeringAction,
				ResultingAction:              protoRule.ResultingAction,
				BoundShifts:                  h.encodeBoundShiftMap(protoRule.BoundShifts),
				UsesComposure:                protoRule.UsesComposure,
				UsesLikeness:                 protoRule.UsesLikeness,
				EncodedShiftMagnitudeChanges: h.encodeStringInt32Map(protoRule.EmotionMagnitudes),
			}
			txErr := tx.Where("rule_id = ?", personalityId).Create(rule)
			if txErr.Error != nil {
				return txErr.Error
			}
			// create rule parts now.
			err := h.CreateRuleParts(rule.RuleID, protoRule.RuleParts, tx)
			if err != nil {
				return err
			}
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = savePersonality(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = savePersonality(tx)
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

func (h *HumanizeDbImpl) GetPersonality(
	personalityIds []string,
	upperTx *gorm.DB,
) ([]*humanize_protobuf.EmotionUpdateRule, error) {
	personality := make([]*humanize_protobuf.EmotionUpdateRule, 0)
	// first get all of the rules
	getPersonality := func(tx *gorm.DB) error {
		var rules []EmotionalBoundRule
		for _, personalityId := range personalityIds {
			var subRules []EmotionalBoundRule
			err := h.mainDB.
				Where("personality_id = ?", personalityId).
				Find(&subRules).
				Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					fmt.Println(err.Error())
				}
				return err
			}
			rules = append(rules, subRules...)
		}
		for _, rule := range rules {
			// loop through all rules, create proto object then get all personalities
			decodedBoundShiftMap, err := h.decodeBoundShiftMap(rule.BoundShifts)
			if err != nil {
				return err
			}
			protoRule := &humanize_protobuf.EmotionUpdateRule{
				RuleId:           rule.RuleID,
				Name:             rule.Name,
				PercentageOfProc: rule.PercentageOfProc,
				TriggeringAction: rule.TriggeringAction,
				ResultingAction:  rule.ResultingAction,
				BoundShifts:      decodedBoundShiftMap,
				UsesComposure:    rule.UsesComposure,
				UsesLikeness:     rule.UsesLikeness,
			}
			personality = append(personality, protoRule)
			// now load rule parts
			ruleParts, err := h.GetRuleParts(rule.RuleID, tx)
			if err != nil {
				return err
			}
			for _, rulePart := range ruleParts {
				protoRule.RuleParts = append(protoRule.RuleParts, rulePart)
			}

		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = getPersonality(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			return getPersonality(tx)
		})
	}
	if err != nil {
		return nil, err
	}
	return personality, nil
}

func (h *HumanizeDbImpl) ListPersonalities() ([]string, error) {
	var arr []string
	err := h.mainDB.Model(Personality{}).Select("personality_id").Find(&arr).Error
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func (h *HumanizeDbImpl) DeletePersonality(
	personalityId string,
	upperTx *gorm.DB,
) error {
	deletePersonalityRule := func(tx *gorm.DB) error {
		err := h.mainDB.Where("personality_id = ?", personalityId).Delete(&EmotionalBoundRule{})
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
