package rule_system

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sort"
	"testserver/conversation_code/prompt_management"
	"testserver/conversation_code/rule_system_utilties"
	"testserver/db"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type PromptRuleSystemManager interface {
	GenerateCharacterPrompt(
		session db.Session,
		memLog *humanize_protobuf.MemoryLog,
		askerName, responderName string,
		npcInformation *db.NpcData,
	) (*humanize_protobuf.Prompt, error)
}

type PromptRuleSystemManagerImpl struct {
	db                    db.HumanizeDB
	emotionalStateManager EmotionalStateManager
}

func NewActuationRuleBasedSystemManager(
	db db.HumanizeDB,
	emotionalStateManager EmotionalStateManager,
) PromptRuleSystemManager {
	return &PromptRuleSystemManagerImpl{
		db:                    db,
		emotionalStateManager: emotionalStateManager,
	}
}

type IdAndDistance struct {
	Id                         string
	DifferenceFromCurrentState float64
}

func remove(slice []humanize_protobuf.PromptSegmentType, s int) []humanize_protobuf.PromptSegmentType {
	return append(slice[:s], slice[s+1:]...)
}

func (a PromptRuleSystemManagerImpl) GenerateCharacterPrompt(
	session db.Session,
	memLog *humanize_protobuf.MemoryLog,
	askerName, responderName string,
	npcInformation *db.NpcData,
) (*humanize_protobuf.Prompt, error) {
	promptSetId := npcInformation.Entity.PromptSetId
	promptSegmentSetId := npcInformation.Entity.PromptSegmentSetId
	log.WithFields(log.Fields{
		"session_id": session.SessionID,
	}).Info("Attempting to generate prompt")
	// first get the set of possible prompts
	idAndStates, err := a.db.GetPromptIdealStatesBySetId(promptSetId, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"session_id": session.SessionID,
		}).Error("could not read possible prompts from database")
		return nil, err
	}
	// then find the current emotional state squared difference
	currentEmotionalState := npcInformation.EmotionalState
	// now get all the ideal square differences for the prompts (map[bound_id]difference from current)
	squareDifferences := make([]IdAndDistance, 0)
	for _, idAndState := range idAndStates {
		squareDifferences = append(squareDifferences, IdAndDistance{
			Id:                         idAndState.Id,
			DifferenceFromCurrentState: CalculateDifferenceBetweenTwoStates(currentEmotionalState, idAndState.State),
		})
	}
	// sort map into descending list
	sort.Slice(squareDifferences, func(i, j int) bool {
		return squareDifferences[i].DifferenceFromCurrentState <
			squareDifferences[j].DifferenceFromCurrentState
	})
	// starting with the prompt closest to the difference
	var selectedPrompt *humanize_protobuf.Prompt
	for _, promptDifference := range squareDifferences {
		log.WithFields(log.Fields{
			"prompt_id": promptDifference.Id,
		}).Info("processing rules for prompt")
		prompt, err := a.db.GetPrompt(promptDifference.Id, nil)
		if err != nil {
			return nil, err
		}
		// Go through the rules
		allRulePartsPassed := true
		for _, rulePart := range prompt.RuleParts {
			// Processing rule part
			log.WithFields(log.Fields{
				"prompt_id":    promptDifference.Id,
				"rule_part_id": rulePart.RulePartId,
			}).Info("assessing rule part")
			perc := rand.Int31n(100)
			if perc >= (100 - rulePart.PercentageOfProc) {
				// find emotional bound
				bound := npcInformation.EmotionalState.EmotionalBounds[rulePart.BoundId]
				if !rule_system_utilties.AssessRulePart(rulePart, bound, memLog) {
					allRulePartsPassed = false
					log.WithFields(log.Fields{
						"prompt_id":          promptDifference.Id,
						"rule_part_id":       rulePart.RulePartId,
						"percentage_of_proc": rulePart.PercentageOfProc,
						"percentage":         perc,
					}).Info("rule part failed")
					break
				}
				log.WithFields(log.Fields{
					"prompt_id":          promptDifference.Id,
					"rule_part_id":       rulePart.RulePartId,
					"percentage_of_proc": rulePart.PercentageOfProc,
					"percentage":         perc,
				}).Info("rule part passed")
			} else {
				log.WithFields(log.Fields{
					"prompt_id":          promptDifference.Id,
					"rule_part_id":       rulePart.RulePartId,
					"percentage_of_proc": rulePart.PercentageOfProc,
					"percentage":         perc,
				}).Info("rule part failed as it did not pass the percentage to proc")
				allRulePartsPassed = false
				break
			}
		}
		if allRulePartsPassed {
			selectedPrompt = prompt
		}
	}
	if selectedPrompt == nil {
		log.WithFields(log.Fields{
			"session_id": session.SessionID,
		}).Error("rule part failed as it did not pass the percentage to proc")
		return nil, errors.New("could not allocate main prompt template")
	}
	promptSegments := make([]*humanize_protobuf.PromptSegment, 0)
	// As long as the prompt has possible prompt segments
	if len(promptSegmentSetId) > 0 {
		// For each needed primer
		for _, neededPrimerType := range selectedPrompt.RequiredPromptSegmentTypes {
			log.WithFields(log.Fields{
				"prompt_id":             selectedPrompt.PromptId,
				"prompt_segment_set_it": promptSegmentSetId,
			}).Info("attempting to get the ideal emotional states for the prompt_segment_set/primer_type combo")
			idAndStates, err :=
				a.db.GetPromptSegmentIdealStateBySetIdAndPrimerType(promptSetId, neededPrimerType.String(), nil)
			if err != nil {
				return nil, err
			}
			log.WithFields(log.Fields{
				"primer_type":           neededPrimerType.String(),
				"prompt_segment_set_id": promptSegmentSetId,
				"possible segments":     len(idAndStates),
			}).Info("possible prompt segments for the primer type")
			// Now calculate prompt segments to see if they should be added
			// To do this we need to order segments by distance squared from current emotional state.
			segmentDistanceSquaredFromEmotionalState := make([]IdAndDistance, 0)
			for _, segment := range idAndStates {
				segmentDistanceSquaredFromEmotionalState =
					append(segmentDistanceSquaredFromEmotionalState, IdAndDistance{
						Id:                         segment.Id,
						DifferenceFromCurrentState: CalculateDifferenceBetweenTwoStates(currentEmotionalState, segment.State),
					})
			}
			sort.SliceStable(segmentDistanceSquaredFromEmotionalState, func(i, j int) bool {
				return segmentDistanceSquaredFromEmotionalState[i].DifferenceFromCurrentState <
					segmentDistanceSquaredFromEmotionalState[j].DifferenceFromCurrentState
			})
			for _, segmentDetails := range segmentDistanceSquaredFromEmotionalState {
				segment, err := a.db.GetPromptSegment(segmentDetails.Id, nil)
				if err != nil {
					return nil, err
				}
				perc := rand.Int31n(100)
				if perc >= (100 - segment.PercentageOfProc) {
					log.WithFields(log.Fields{
						"prompt_id":             selectedPrompt.PromptId,
						"percentage_of_proc":    segment.PercentageOfProc,
						"prompt_segment_type":   segment.Type,
						"segment_id":            segment.PromptSegmentId,
						"percentage":            perc,
						"prompt_segment_set_id": promptSegmentSetId,
					}).Info("prompt segment has passed percentage check, now " +
						"attempting to process rule parts")
					allRulePartsPassed := true
					for _, rulePart := range segment.RuleParts {
						perc := rand.Int31n(100)
						if perc >= (100 - segment.PercentageOfProc) {
							log.WithFields(log.Fields{
								"prompt_id":             selectedPrompt.PromptId,
								"percentage_of_proc":    rulePart.PercentageOfProc,
								"prompt_segment_type":   segment.Type,
								"segment_id":            segment.PromptSegmentId,
								"percentage":            perc,
								"prompt_segment_set_it": promptSegmentSetId,
								"segment_rule_part_id":  rulePart.RulePartId,
							}).Info("processing rule part")
							bound := npcInformation.EmotionalState.EmotionalBounds[rulePart.BoundId]
							success := rule_system_utilties.AssessRulePart(rulePart, bound, memLog)
							if success {
								log.WithFields(log.Fields{
									"prompt_id":             selectedPrompt.PromptId,
									"percentage_of_proc":    rulePart.PercentageOfProc,
									"prompt_segment_type":   segment.Type,
									"segment_id":            segment.PromptSegmentId,
									"percentage":            perc,
									"prompt_segment_set_it": promptSegmentSetId,
									"segment_rule_part_id":  rulePart.RulePartId,
								}).Info("rule part passed")
							} else {
								log.WithFields(log.Fields{
									"prompt_id":             selectedPrompt.PromptId,
									"percentage_of_proc":    rulePart.PercentageOfProc,
									"prompt_segment_type":   segment.Type,
									"segment_id":            segment.PromptSegmentId,
									"percentage":            perc,
									"prompt_segment_set_it": promptSegmentSetId,
									"segment_rule_part_id":  rulePart.RulePartId,
								}).Info("rule part failed")
								allRulePartsPassed = false
								break
							}
						} else {
							log.WithFields(log.Fields{
								"prompt_id":             selectedPrompt.PromptId,
								"percentage_of_proc":    rulePart.PercentageOfProc,
								"prompt_segment_type":   segment.Type,
								"segment_id":            segment.PromptSegmentId,
								"percentage":            perc,
								"prompt_segment_set_it": promptSegmentSetId,
								"segment_rule_part_id":  rulePart.RulePartId,
							}).Info("prompt segment rule part did not pass")
						}
					}
					if allRulePartsPassed {
						log.WithFields(log.Fields{
							"prompt_id":             selectedPrompt.PromptId,
							"percentage_of_proc":    segment.PercentageOfProc,
							"prompt_segment_type":   segment.Type,
							"segment_id":            segment.PromptSegmentId,
							"percentage":            perc,
							"prompt_segment_set_it": promptSegmentSetId,
						}).Info("all rule parts have passed")
					}
				} else {
					log.WithFields(log.Fields{
						"prompt_id":             selectedPrompt.PromptId,
						"percentage_of_proc":    segment.PercentageOfProc,
						"percentage":            perc,
						"prompt_segment_set_it": promptSegmentSetId,
					}).Info("prompt segment has failed percentage check")
				}
			}
		}
	}
	finalPrompt, err := prompt_management.GeneratePromptFromList(
		askerName, responderName,
		selectedPrompt.PromptText, promptSegments,
		memLog, npcInformation.EmotionalState,
	)
	if err != nil {
		return nil, err
	}
	return &humanize_protobuf.Prompt{
		PromptText: finalPrompt,
	}, nil
}
