package rule_system_utilties

import (
	"math"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

func AssessRulePart(
	rulePart *humanize_protobuf.RulePart,
	bound *humanize_protobuf.EmotionalBound,
	memLog *humanize_protobuf.MemoryLog,
) bool {
	switch rulePart.Type {
	case humanize_protobuf.RulePart_IS_POSITIVE:
		return bound.CurrentPercentage >= 0
	case humanize_protobuf.RulePart_IS_NEGATIVE:
		return bound.CurrentPercentage <= 0
	case humanize_protobuf.RulePart_IS_NEGATIVE_BELOW_X_PERCENT:
		if bound.CurrentPercentage > 0 {
			return false
		}
		return bound.CurrentPercentage > rulePart.BoundPercentage
	case humanize_protobuf.RulePart_IS_POSITIVE_ABOVE_X_PERCENT:
		if bound.CurrentPercentage < 0 {
			return false
		}
		return int32(math.Abs(float64(bound.CurrentPercentage))) > rulePart.BoundPercentage
	case humanize_protobuf.RulePart_IS_BELOW_X_PERCENT:
		return bound.CurrentPercentage <= rulePart.BoundPercentage
	case humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT:
		return bound.CurrentPercentage > rulePart.BoundPercentage
	case humanize_protobuf.RulePart_EMOTIONAL_SEQUENCE:
		//	emotionTrain := rulePart.ConversationEmotionTrain
		//	if memLog != nil {
		//		if len(memLog.Log) >= len(emotionTrain) {
		//			for i := 0; i < len(emotionTrain); i++ {
		//				actualEmotion := memLog.Log[len(memLog.Log)-len(emotionTrain)+i].ResponseEmotionClassification
		//				if actualEmotion != emotionTrain[i] {
		//					return false
		//				}
		//			}
		//			return true
		//		}
		//	}
		return false
	}
	return false
}
