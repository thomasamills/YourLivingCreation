package rule_system

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"testserver/conversation_code/npc_data"
	"testserver/conversation_code/rule_system_utilties"
	"testserver/db"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type EmotionalStateManager interface {
	UpdateEmotionalState(
		entityId string,
		boundShifts map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType,
		boundMagnitudes map[string]humanize_protobuf.EmotionShiftMagnitude,
		usesComposure bool, composure int32,
	) error
	ProcessEmotionalStateRules(
		npc npc_data.NpcData, triggers []string,
	) (*humanize_protobuf.EmotionalState, []string, error)
	GetEmotionalState(entityId string) (*humanize_protobuf.EmotionalState, error)

	LoadEmotionalStateFromPreset(
		presetStateId, entityId string,
	) (*humanize_protobuf.EmotionalState, error)
}

type EmotionalStateManagerImpl struct {
	db db.HumanizeDB
}

func NewEmotionalStateManager(db db.HumanizeDB) EmotionalStateManager {
	return &EmotionalStateManagerImpl{
		db: db,
	}
}

func (e *EmotionalStateManagerImpl) GetEmotionalState(
	entityId string,
) (*humanize_protobuf.EmotionalState, error) {
	return e.db.GetEmotionalState(entityId, nil)
}

func (e *EmotionalStateManagerImpl) UpdateEmotionalState(
	entityId string,
	boundShifts map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType,
	boundMagnitudes map[string]humanize_protobuf.EmotionShiftMagnitude,
	usesComposure bool, composure int32,
) error {
	emotionalState, err := e.db.GetEmotionalState(entityId, nil)
	if err != nil {
		return err
	}
	// First update the bound magnitudes
	for key, value := range boundMagnitudes {
		emotionalState.EmotionalBounds[key].EmotionShiftMagnitude = value
	}
	getInitialShiftValue := func(
		shiftType humanize_protobuf.EmotionUpdateRule_BoundShiftType,
	) float32 {
		switch shiftType {
		case humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW:
			return -20 - float32(rand.Int31n(10))
		case humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM:
			return -40 - float32(rand.Int31n(10))
		case humanize_protobuf.EmotionUpdateRule_NEGATIVE_HIGH:
			return -60 - float32(rand.Int31n(10))
		case humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW:
			return 20 + float32(rand.Int31n(10))
		case humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM:
			return 40 + float32(rand.Int31n(10))
		case humanize_protobuf.EmotionUpdateRule_POSITIVE_HIGH:
			return 60 + float32(rand.Int31n(10))
		}
		return 0
	}
	getRandomShiftMagnitudeValue := func(
		magnitudeShift humanize_protobuf.EmotionShiftMagnitude,
	) float32 {
		switch magnitudeShift {
		case humanize_protobuf.EmotionShiftMagnitude_BOTTLING_UP:
			return 0.25 + (float32(rand.Int31n(10) / 100))
		case humanize_protobuf.EmotionShiftMagnitude_NEUTRAL:
			return 0.5 + (float32(rand.Int31n(10) / 100))
		case humanize_protobuf.EmotionShiftMagnitude_SHORT_FUSE:
			return 0.9 + (float32(rand.Int31n(10) / 100))
		}
		return 0
	}
	// Create map to store the final bound shifts that will be applied.
	finalBoundShifts := make(map[string]float32, 0)
	for boundId, boundShiftType := range boundShifts {
		finalBoundShifts[boundId] = getInitialShiftValue(boundShiftType)
	}
	// Update bound shifts via composure before application
	if usesComposure {
		for key, shift := range finalBoundShifts {
			composureNormalization := (100 - float32(composure)) / 100
			finalBoundShifts[key] = shift * composureNormalization
		}
	}
	// Update bound shifts via shift magnitudes before application
	if len(boundMagnitudes) > 0 {
		for boundId, bound := range emotionalState.EmotionalBounds {
			finalBoundShifts[boundId] =
				finalBoundShifts[boundId] *
					getRandomShiftMagnitudeValue(bound.EmotionShiftMagnitude)
		}
	}
	// Apply bound shifts
	for key, value := range finalBoundShifts {
		log.WithFields(log.Fields{
			key: value,
		}).Info("Updating emotional bound")
		bound := emotionalState.EmotionalBounds[key]
		bound.CurrentPercentage += int32(value)
		if bound.CurrentPercentage < -50 {
			bound.CurrentPercentage = -50
		}
		if bound.CurrentPercentage > 50 {
			bound.CurrentPercentage = 50
		}
	}
	// TODO Now smooth the emotional state.

	// Update bound shifts + new magnitudes to the database.
	for _, bound := range emotionalState.EmotionalBounds {
		err := e.db.UpdateEmotionalBoundPercentage(bound)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *EmotionalStateManagerImpl) ProcessEmotionalStateRules(
	npc npc_data.NpcData, triggers []string,
) (*humanize_protobuf.EmotionalState, []string, error) {
	triggerString := ""
	for _, trigger := range triggers {
		triggerString += trigger + " "
	}
	log.WithFields(log.Fields{
		"entity_id": npc.Entity.EntityId,
		"triggers":  triggerString,
	}).Info("Processing emotional state rules")
	resultingActions := make([]string, 0)
	for _, rule := range npc.Personality {
		for _, trigger := range triggers {
			if rule.TriggeringAction == trigger {
				log.WithFields(log.Fields{
					"triggering_action": rule.TriggeringAction,
					"rule_name":         rule.Name,
				}).Info("rule triggered for assessment")
				perc := rand.Int31n(100)
				if perc >= (100 - rule.PercentageOfProc) {
					log.WithFields(log.Fields{
						"triggering_action":     rule.TriggeringAction,
						"rule_name":             rule.Name,
						"percentage_to_trigger": rule.PercentageOfProc,
						"percentage_shot":       perc,
					}).Info("rule passed percentage of proc ")
					allRulesPassed := true
					for _, rulePart := range rule.RuleParts {
						log.WithFields(log.Fields{
							"triggering_action": rule.TriggeringAction,
							"rule_name":         rule.Name,
							"rule_part_type":    rulePart.Type.String(),
							"bound_percentage":  rulePart.BoundPercentage,
							"bound_id":          rulePart.BoundId,
						}).Info("assessing rule part")
						rulePartPerc := rand.Int31n(100)
						if rulePartPerc >= (100 - rulePart.PercentageOfProc) {
							log.WithFields(log.Fields{
								"triggering_action":     rule.TriggeringAction,
								"rule_name":             rule.Name,
								"percentage_to_trigger": rule.PercentageOfProc,
								"percentage_shot":       perc,
								"rule_part_type":        rulePart.Type.String(),
								"bound_percentage":      rulePart.BoundPercentage,
								"bound_id":              rulePart.BoundId,
							}).Info("rule passed percentage of proc ")
							bound := npc.EmotionalState.EmotionalBounds[rulePart.BoundId]
							if bound == nil {
								log.WithFields(log.Fields{
									"triggering_action":     rule.TriggeringAction,
									"rule_name":             rule.Name,
									"percentage_to_trigger": rule.PercentageOfProc,
									"percentage_shot":       perc,
									"rule_part_type":        rulePart.Type.String(),
									"bound_percentage":      rulePart.BoundPercentage,
									"bound_id":              rulePart.BoundId,
								}).Error("could not find emotional bound ")
								return nil, nil,
									errors.New(
										fmt.Sprintf(
											"could not retrieve emotional bound to assess rule part"+
												" for rule part %s", rulePart.RuleId,
										),
									)
							}
							rulePartSuccess := rule_system_utilties.AssessRulePart(rulePart, bound, nil)
							if !rulePartSuccess {
								log.WithFields(log.Fields{
									"triggering_action":     rule.TriggeringAction,
									"rule_name":             rule.Name,
									"percentage_to_trigger": rule.PercentageOfProc,
									"percentage_shot":       perc,
									"rule_part_type":        rulePart.Type.String(),
									"bound_percentage":      rulePart.BoundPercentage,
									"bound_id":              rulePart.BoundId,
								}).Info("rule failed ")
								allRulesPassed = false
							}
							log.WithFields(log.Fields{
								"triggering_action":     rule.TriggeringAction,
								"rule_name":             rule.Name,
								"percentage_to_trigger": rule.PercentageOfProc,
								"percentage_shot":       perc,
								"rule_part_type":        rulePart.Type.String(),
								"bound_percentage":      rulePart.BoundPercentage,
								"bound_id":              rulePart.BoundId,
							}).Info("rule passed")
						}
					}
					if allRulesPassed {
						err := e.UpdateEmotionalState(
							npc.Entity.EntityId, rule.BoundShifts,
							rule.EmotionMagnitudes,
							rule.UsesComposure, npc.EmotionalState.Composure)
						if err != nil {
							return nil, nil, err
						}
						if len(rule.ResultingAction) > 0 {
							resultingActions = append(resultingActions, rule.ResultingAction)
						}
					}
				}
			}
		}
	}
	finalState, err := e.GetEmotionalState(npc.Entity.EntityId)
	if err != nil {
		return nil, nil, err
	}
	return finalState, resultingActions, nil
}

// LoadEmotionalStateFromPreset Copies emotional state from a pre-saved one
func (e *EmotionalStateManagerImpl) LoadEmotionalStateFromPreset(
	presetStateId, entityId string,
) (*humanize_protobuf.EmotionalState, error) {
	emotionalState, err := e.db.CopyEmotionalState(presetStateId, entityId)
	if err != nil {
		return nil, err
	}
	if !emotionalState.IsPreset {
		return nil, errors.New("cannot load non preset emotional state")
	}
	emotionalState.IsPreset = false
	return emotionalState, nil
}

func CalculateDifferenceBetweenTwoStates(x, y *humanize_protobuf.EmotionalState) float64 {
	difference := float64(0)
	for _, bound := range x.EmotionalBounds {
		correspondingBound := y.EmotionalBounds[bound.BoundId]
		xPer := bound.CurrentPercentage
		yPer := correspondingBound.CurrentPercentage
		difference += math.Abs(float64(xPer - yPer))
	}
	return difference
}
