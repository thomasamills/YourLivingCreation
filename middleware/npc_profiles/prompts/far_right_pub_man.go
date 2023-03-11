package prompts

import (
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

var FarRightPubManPromptNormal = `{{.StoryBackground}}
${{.MemLog}}
${{.MotivationalPrimer}}
${{.EmotionalPrimer}}
${{.IdeologyPrimer}}
${{.ReligionPrimer}}
${{.EmotionalPrimer}}
${{.Topic}}.`

var FarRightPubManNormal = &humanize_protobuf.Prompt{
	PromptId:    "FAR_RIGHT_PUB_MAN_PROMPT",
	PromptSetId: "FAR_RIGHT_PUB_MAN_PROMPT_SET",
	PromptName:  "FAR_RIGHT_PUB_MAN_PROMPT",
	PromptText:  FarRightPubManPromptNormal,
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
	RequiredPromptSegmentTypes: []humanize_protobuf.PromptSegmentType{
		humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_STORY_BACKGROUND,
		humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
		humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
		humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_RELIGION_PRIMER,
		humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_EMOTIONAL_PRIMER,
		humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
	},
}
