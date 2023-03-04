package rule_system

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"sort"
	"testserver/db"
	npcdata "testserver/npcs"
	"testserver/prompt_management"
	"testserver/rule_system_utilties"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type PromptRuleSystemManager interface {
	GenerateCharacterPrompt(
		session db.Session,
		memLog *humanize_protobuf.MemoryLog,
		askerName, responderName string,
		npcInformation *npcdata.NpcData,
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
	PromptId                   string
	DifferenceFromCurrentState float64
}

func remove(slice []humanize_protobuf.PromptSegmentType, s int) []humanize_protobuf.PromptSegmentType {
	return append(slice[:s], slice[s+1:]...)
}

func (a PromptRuleSystemManagerImpl) GenerateCharacterPrompt(
	session db.Session,
	memLog *humanize_protobuf.MemoryLog,
	askerName, responderName string,
	npcInformation *npcdata.NpcData,
) (*humanize_protobuf.Prompt, error) {
	promptSetId := npcInformation.Entity.PromptSetId
	promptSegmentSetId := npcInformation.Entity.PromptSegmentSetId
	log.WithFields(log.Fields{
		"session_id": session.SessionID,
	}).Info("Attempting to generate prompt")
	// first get the set of possible prompts
	prompts, err := a.db.GetPromptSet(promptSetId, nil)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{
		"session_id":    session.SessionID,
		"prompt_set_id": promptSetId,
	}).Info(fmt.Sprintf("%d possible prompts", len(prompts)))
	if len(prompts) == 0 {
		return nil, errors.New("no possible prompts")
	}
	// then find the current emotional state squared difference
	currentEmotionalStateSquaredDifference :=
		rule_system_utilties.CalculateEmotionalStateSquaredDifference(npcInformation.EmotionalState)
	// now get all the ideal square differences for the prompts (map[bound_id]difference from current)
	squareDifferences := make([]IdAndDistance, 0)
	for _, prompt := range prompts {
		squareDifferences = append(squareDifferences, IdAndDistance{
			PromptId: prompt.PromptId,
			DifferenceFromCurrentState: math.Abs(
				float64(prompt.EmotionalStateSquaredDifference - currentEmotionalStateSquaredDifference),
			),
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
			"prompt_id": promptDifference.PromptId,
		}).Info("processing rules for prompt")
		prompt, err := a.db.GetPrompt(promptDifference.PromptId, nil)
		if err != nil {
			return nil, err
		}
		// Go through the rules
		allRulePartsPassed := true
		for _, rulePart := range prompt.RuleParts {
			// Processing rule part
			log.WithFields(log.Fields{
				"prompt_id":    promptDifference.PromptId,
				"rule_part_id": rulePart.RulePartId,
			}).Info("assessing rule part")
			perc := rand.Int31n(100)
			if perc >= (100 - rulePart.PercentageOfProc) {
				// find emotional bound
				bound := npcInformation.EmotionalState.EmotionalBounds[rulePart.BoundId]
				if !rule_system_utilties.AssessRulePart(rulePart, bound, memLog) {
					allRulePartsPassed = false
					log.WithFields(log.Fields{
						"prompt_id":          promptDifference.PromptId,
						"rule_part_id":       rulePart.RulePartId,
						"percentage_of_proc": rulePart.PercentageOfProc,
						"percentage":         perc,
					}).Info("rule part failed")
					break
				}
				log.WithFields(log.Fields{
					"prompt_id":          promptDifference.PromptId,
					"rule_part_id":       rulePart.RulePartId,
					"percentage_of_proc": rulePart.PercentageOfProc,
					"percentage":         perc,
				}).Info("rule part passed")
			} else {
				log.WithFields(log.Fields{
					"prompt_id":          promptDifference.PromptId,
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
		log.WithFields(log.Fields{
			"prompt_id":             selectedPrompt.PromptId,
			"prompt_segment_set_it": promptSegmentSetId,
		}).Info("attempting to get the prompt segment set")
		promptSegmentSet, err := a.db.GetPromptSegmentSet(promptSegmentSetId, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"prompt_id":             selectedPrompt.PromptId,
				"prompt_segment_set_it": promptSegmentSetId,
				"error_message":         err.Error(),
			}).Error("error while attempting to get prompt segment set")
			return nil, err
		}
		log.WithFields(log.Fields{
			"prompt_id":             selectedPrompt.PromptId,
			"prompt_segment_set_it": promptSegmentSetId,
		}).Info("processing all possible prompt segments")
		// Now calculate prompt segments to see if they should be added
		// To do this we need to order segments by distance squared from current emotional state.
		segmentDistanceSquaredFromEmotionalState := make([]IdAndDistance, 0)
		for _, segment := range promptSegmentSet {
			segmentDistanceSquaredFromEmotionalState =
				append(segmentDistanceSquaredFromEmotionalState, IdAndDistance{
					PromptId:                   segment.PromptSegmentId,
					DifferenceFromCurrentState: float64(currentEmotionalStateSquaredDifference),
				})
		}
		sort.SliceStable(segmentDistanceSquaredFromEmotionalState, func(i, j int) bool {
			return segmentDistanceSquaredFromEmotionalState[i].DifferenceFromCurrentState <
				segmentDistanceSquaredFromEmotionalState[j].DifferenceFromCurrentState
		})
		// to store which prompts are needed
		requiredSegmentsForPrompt := make([]humanize_protobuf.PromptSegmentType, 0)
		for _, neededPrompt := range selectedPrompt.RequiredPromptSegmentTypes {
			requiredSegmentsForPrompt = append(requiredSegmentsForPrompt, neededPrompt)
		}
		for _, segment := range promptSegmentSet {
			// First check if the prompt can hold another one of these primers
			log.WithFields(log.Fields{
				"session_id":        session.SessionID,
				"prompt_id":         selectedPrompt.PromptId,
				"prompt_segment_id": segment.PromptSegmentId,
			}).Info("processing prompt segment")
			containsSegment := false
			for _, segmentType := range requiredSegmentsForPrompt {
				if segment.Type == segmentType {
					containsSegment = true
					break
				}
			}
			if containsSegment {
				// calculate percentage of proc
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
						promptSegments = append(promptSegments, segment)
						// now we need to remove this from the needed prompts
						for i, neededSegmentType := range requiredSegmentsForPrompt {
							if segment.Type == neededSegmentType {
								remove(requiredSegmentsForPrompt, i)
							}
						}
						log.WithFields(log.Fields{
							"prompt_id":             selectedPrompt.PromptId,
							"percentage_of_proc":    segment.PercentageOfProc,
							"prompt_segment_type":   segment.Type,
							"segment_id":            segment.PromptSegmentId,
							"percentage":            perc,
							"prompt_segment_set_it": promptSegmentSetId,
						}).Info("no remaining segments")
						if len(requiredSegmentsForPrompt) == 0 {
							break
						}
					}
				} else {
					log.WithFields(log.Fields{
						"prompt_id":             selectedPrompt.PromptId,
						"percentage_of_proc":    segment.PercentageOfProc,
						"percentage":            perc,
						"prompt_segment_set_it": promptSegmentSetId,
					}).Info("prompt segment has failed percentage check")
				}
			} else {
				log.WithFields(log.Fields{
					"session_id":        session.SessionID,
					"prompt_id":         selectedPrompt.PromptId,
					"prompt_segment_id": segment.PromptSegmentId,
					"segment_type":      segment.Type,
				}).Info("Segment type no longer required")
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
