package rule_system

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"testserver/conversation_code/rule_system_utilties"
	"testserver/db"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type ActuationRuleManager interface {
	ProcessActuationRules(
		entityId string, session db.Session,
	) error
}

type ActuationRuleManagerImpl struct {
	db db.HumanizeDB
}

func (arm *ActuationRuleManagerImpl) ProcessActuationRules(
	entityId string,
	session *db.Session,
) error {
	log.WithFields(log.Fields{
		"entity_id": entityId,
	}).Info("Processing actuation rules")
	entity, err := arm.db.GetEntity(entityId)
	if err != nil {
		return err
	}
	actuationRuleSet, err := arm.db.GetActuationRuleSet(entity.PersonalityId, nil)
	if err != nil {
		return err
	}
	emotionalState, err := arm.db.GetEmotionalState(entityId, nil)
	if err != nil {
		return err
	}
	for _, rule := range actuationRuleSet {
		perc := rand.Int31n(100)
		if perc >= (100 - rule.PercentageOfProc) {
			log.WithFields(log.Fields{
				"rule_name":             rule.RuleName,
				"percentage_to_trigger": rule.PercentageOfProc,
				"percentage_shot":       perc,
			}).Info("rule passed percentage of proc ")
			allRulePartsPassed := true
			for _, rulePart := range rule.RuleParts {
				log.WithFields(log.Fields{
					"rule_name":        rule.RuleName,
					"rule_part_type":   rulePart.Type.String(),
					"bound_percentage": rulePart.BoundPercentage,
					"bound_id":         rulePart.BoundId,
				}).Info("assessing rule part")
				rulePartPerc := rand.Int31n(100)
				if rulePartPerc >= (100 - rulePart.PercentageOfProc) {
					log.WithFields(log.Fields{
						"rule_name":             rule.RuleName,
						"percentage_to_trigger": rule.PercentageOfProc,
						"percentage_shot":       perc,
						"rule_part_type":        rulePart.Type.String(),
						"bound_percentage":      rulePart.BoundPercentage,
						"bound_id":              rulePart.BoundId,
					}).Info("rule passed percentage of proc ")
					bound := emotionalState.EmotionalBounds[rulePart.BoundId]
					if bound == nil {
						log.WithFields(log.Fields{
							"rule_name":             rule.RuleName,
							"percentage_to_trigger": rule.PercentageOfProc,
							"percentage_shot":       perc,
							"rule_part_type":        rulePart.Type.String(),
							"bound_percentage":      rulePart.BoundPercentage,
							"bound_id":              rulePart.BoundId,
						}).Error("could not find emotional bound ")
						return errors.New(
							fmt.Sprintf(
								"could not retrieve emotional bound to assess rule part"+
									" for rule part %s", rulePart.RuleId,
							),
						)
					}
					rulePartSuccess := rule_system_utilties.AssessRulePart(rulePart, bound, nil)
					if !rulePartSuccess {
						log.WithFields(log.Fields{
							"rule_name":             rule.RuleName,
							"percentage_to_trigger": rule.PercentageOfProc,
							"percentage_shot":       perc,
							"rule_part_type":        rulePart.Type.String(),
							"bound_percentage":      rulePart.BoundPercentage,
							"bound_id":              rulePart.BoundId,
						}).Info("rule failed")
						allRulePartsPassed = false
					}
					log.WithFields(log.Fields{
						"rule_name":             rule.RuleName,
						"percentage_to_trigger": rule.PercentageOfProc,
						"percentage_shot":       perc,
						"rule_part_type":        rulePart.Type.String(),
						"bound_percentage":      rulePart.BoundPercentage,
						"bound_id":              rulePart.BoundId,
					}).Info("rule passed")
				}
			}
			if allRulePartsPassed {
				err := arm.ApplyRule(rule, session)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (arm *ActuationRuleManagerImpl) ApplyRule(
	rule *humanize_protobuf.ActuationRule, session *db.Session,
) error {
	return nil
}
