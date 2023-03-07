package emotional_primers

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

var DefaultEmotionalPrimerNormal = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_EMOTIONAL_PRIMER.String(),
	PromptSegmentId:    "DEFAULT_EMOTIONAL_PRIMER",
	PromptSegmentSetId: "DEFAULT_EMOTIONAL_PRIMER_SET_ID",
	Message:            "{{.ResponderName}} is currently feeling {{.Synonym}}",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_EMOTIONAL_PRIMER,
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 0,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 0,
			},
			"agitated_patient": {
				CurrentPercentage: 0,
			},
			"hate_love": {
				CurrentPercentage: 0,
			},
		},
	},
	// No rule parts is on normal.
}
