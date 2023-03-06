package db

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

func (h *HumanizeDbImpl) CreateEmotionalBoundRule(protoRule *humanize_protobuf.EmotionUpdateRule) (string, error) {
	rule := &EmotionalBoundRule{
		RuleID:                       h.idGen.GetULIDfromtime(),
		PersonalityID:                protoRule.PersonalityId,
		Name:                         protoRule.Name,
		PercentageOfProc:             protoRule.PercentageOfProc,
		TriggeringAction:             protoRule.TriggeringAction,
		ResultingAction:              protoRule.ResultingAction,
		BoundShifts:                  h.encodeBoundShiftMap(protoRule.BoundShifts),
		UsesComposure:                protoRule.UsesComposure,
		UsesLikeness:                 protoRule.UsesLikeness,
		EncodedShiftMagnitudeChanges: h.encodeStringInt32Map(protoRule.EmotionMagnitudes),
	}
	err := h.mainDB.Create(rule)
	if err != nil {
		if err.Error != nil {
			return "", err.Error
		}
	}
	return rule.RuleID, nil
}

func (h *HumanizeDbImpl) CreateEmotionalBoundRulePart(
	protoRulePart *humanize_protobuf.RulePart,
) (string, error) {
	rulePart := &RulePart{
		RulePartId:       h.idGen.GetULIDfromtime(),
		RuleID:           protoRulePart.RulePartId,
		PercentageOfProc: protoRulePart.PercentageOfProc,
		RuleType:         int32(protoRulePart.Type),
		BoundPercentage:  protoRulePart.BoundPercentage,
		BoundId:          protoRulePart.BoundId,
	}
	err := h.mainDB.Create(rulePart)
	if err != nil {
		if err.Error != nil {
			return "", err.Error
		}
	}
	return rulePart.RulePartId, nil
}
