package needs

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

var FoodPromptSegmentNormal = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_NORMAL",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "Food - You may need to have a balanced diet and adequate nutrition to maintain Your physical health. " +
		"You can get this by eating a variety of fruits, vegetables, whole grains, lean proteins, and healthy fats.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
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

var FoodPromptSegmentHappy = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_HAPPY",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You can enjoy your happiness while still prioritizing her physical health through a balanced diet. " +
		"You can celebrate with a healthy meal that includes her favorite fruits, vegetables, and lean proteins.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be above 20 happy",
			BoundPercentage: 20,
			BoundId:         "sad_happy",
			Type:            humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 20,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 5,
			},
			"agitated_patient": {
				CurrentPercentage: 5,
			},
			"hate_love": {
				CurrentPercentage: 5,
			},
		},
	},
}

var FoodPromptSegmentVeryHappy = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_VERY_HAPPY",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You can continue to prioritize her physical health even in moments of extreme happiness." +
		" You can indulge in your favorite healthy foods and use her positive energy to motivate herself to maintain a balanced die",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be above 35 happy",
			BoundPercentage: 35,
			BoundId:         "sad_happy",
			Type:            humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 35,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 10,
			},
			"agitated_patient": {
				CurrentPercentage: 10,
			},
			"hate_love": {
				CurrentPercentage: 10,
			},
		},
	},
}

var FoodPromptSegmentSad = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_SAD",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "Even when feeling sad, You may need to prioritize her physical health by ensuring you eat a balanced diet. " +
		"You can reach out to a friend for support or try to find joy in preparing a healthy and delicious meal.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be below 20 sad",
			BoundPercentage: -20,
			BoundId:         "sad_happy",
			Type:            humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: -20,
			},
			"burned_out_ambitious": {
				CurrentPercentage: -5,
			},
			"agitated_patient": {
				CurrentPercentage: -5,
			},
			"hate_love": {
				CurrentPercentage: -5,
			},
		},
	},
}

var FoodPromptSegmentVerySad = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_VERY_SAD",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "When feeling very sad, You may struggle to find the motivation to take care of herself. However," +
		" she can remind herself that proper nutrition is an essential part of self-care and try to focus on " +
		"easy-to-prepare nutritious meals",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be below 35 sad",
			BoundPercentage: 35,
			BoundId:         "sad_happy",
			Type:            humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: -35,
			},
			"burned_out_ambitious": {
				CurrentPercentage: -10,
			},
			"agitated_patient": {
				CurrentPercentage: -10,
			},
			"hate_love": {
				CurrentPercentage: -10,
			},
		},
	},
}

var FoodPromptSegmentPatient = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_PATIENT",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You can recognize that taking care of your physical health through a balanced diet is a " +
		"long-term commitment that requires patience. She can make a plan to gradually incorporate " +
		"more nutritious foods into her diet and celebrate small victories along the way",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be above 20 patient",
			BoundPercentage: 20,
			BoundId:         "agitated_patient",
			Type:            humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 5,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 5,
			},
			"agitated_patient": {
				CurrentPercentage: 20,
			},
			"hate_love": {
				CurrentPercentage: 5,
			},
		},
	},
}

var FoodPromptSegmentVeryPatient = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_VERY_PATIENT",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You can continue to be patient with herself as you work towards maintaining a balanced diet " +
		"and adequate nutrition. You can research healthy recipes and try new foods" +
		" to keep herself motivated and engaged in the process",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be above 35 patient",
			BoundPercentage: 35,
			BoundId:         "agitated_patient",
			Type:            humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 10,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 10,
			},
			"agitated_patient": {
				CurrentPercentage: 35,
			},
			"hate_love": {
				CurrentPercentage: 10,
			},
		},
	},
}

var FoodPromptSegmentAgitated = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_AGITATED",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You may need to have a balanced diet and adequate nutrition to maintain your physical health," +
		" even when she's angry. You can take a break and prepare herself a nutritious meal to" +
		" help calm you down and support her body",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be below -20 agitated",
			BoundPercentage: -20,
			BoundId:         "agitated_patient",
			Type:            humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: -5,
			},
			"burned_out_ambitious": {
				CurrentPercentage: -5,
			},
			"agitated_patient": {
				CurrentPercentage: -20,
			},
			"hate_love": {
				CurrentPercentage: -5,
			},
		},
	},
}

var FoodPromptSegmentVeryAgitated = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_VERY_AGITATED",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "Despite feeling very angry, You may need to recognize that her body still needs proper nutrition to function." +
		" You can take a few deep breaths and try to find a healthy outlet for your anger," +
		" such as going for a run, before eating a balanced meal",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be below -35 agitated",
			BoundPercentage: -35,
			BoundId:         "agitated_patient",
			Type:            humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: -10,
			},
			"burned_out_ambitious": {
				CurrentPercentage: -10,
			},
			"agitated_patient": {
				CurrentPercentage: -35,
			},
			"hate_love": {
				CurrentPercentage: -10,
			},
		},
	},
}

var FoodPromptSegmentInLove = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_IN_LOVE",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You can share her commitment to a balanced diet with her partner and use cooking together " +
		"as a bonding activity. She can try new healthy recipes and find joy " +
		"in supporting each other's physical health.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be above 20 love",
			BoundPercentage: 20,
			BoundId:         "hate_love",
			Type:            humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 5,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 5,
			},
			"agitated_patient": {
				CurrentPercentage: 5,
			},
			"hate_love": {
				CurrentPercentage: 20,
			},
		},
	},
}

var FoodPromptSegmentMuchInLove = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_MUCH_IN_LOVE",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: " You can prioritize your own physical health while also supporting your partner's health by" +
		" preparing nutritious meals together. She can find balance in her relationship " +
		"and prioritize taking care of herself",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 5,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 10,
			},
			"agitated_patient": {
				CurrentPercentage: 10,
			},
			"hate_love": {
				CurrentPercentage: 35,
			},
		},
	},
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be above 35 love",
			BoundPercentage: 35,
			BoundId:         "hate_love",
			Type:            humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
		},
	},
}

var FoodPromptSegmentHatrid = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_HATE",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "When feeling hate-filled, You may struggle to find the motivation to take care of her physical health. " +
		"However, you can remind yourself that proper nutrition is an essential part of" +
		" self-care and try to focus on easy-to-prepare nutritious meals",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be below -20 hate",
			BoundPercentage: -20,
			BoundId:         "hate_love",
			Type:            humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 10,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 10,
			},
			"agitated_patient": {
				CurrentPercentage: 10,
			},
			"hate_love": {
				CurrentPercentage: 35,
			},
		},
	},
}

var FoodPromptSegmentMuchHatrid = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_MUCH_HATE",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "Despite feeling very hate-filled, You may need to recognize that her body still needs proper nutrition" +
		" to function. She can take a few deep breaths and try to find a healthy outlet for her" +
		" emotions before eating a balanced meal",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be below -35 hate",
			BoundPercentage: -35,
			BoundId:         "hate_love",
			Type:            humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: -10,
			},
			"burned_out_ambitious": {
				CurrentPercentage: -10,
			},
			"agitated_patient": {
				CurrentPercentage: -10,
			},
			"hate_love": {
				CurrentPercentage: -35,
			},
		},
	},
}

var FoodPromptSegmentAmbitious = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_AMBITIOUS",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You can use your ambition to fuel your commitment to maintaining a balanced diet and adequate " +
		"nutrition. You can set goals for herself and track your progress to stay motivated",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be above 20 ambitious",
			BoundPercentage: 20,
			BoundId:         "burned_out_ambitious",
			Type:            humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 5,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 20,
			},
			"agitated_patient": {
				CurrentPercentage: 5,
			},
			"hate_love": {
				CurrentPercentage: 5,
			},
		},
	},
}

var FoodPromptSegmentVeryAmbitious = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_VERY_AMBITIOUS",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You can continue to challenge yourself to maintain a balanced diet and adequate nutrition " +
		"while also pursuing her ambitious goals. You can prioritize her physical health to support your energy and focus.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be above 35 ambitious",
			BoundPercentage: 35,
			BoundId:         "burned_out_ambitious",
			Type:            humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: 10,
			},
			"burned_out_ambitious": {
				CurrentPercentage: 35,
			},
			"agitated_patient": {
				CurrentPercentage: 10,
			},
			"hate_love": {
				CurrentPercentage: 10,
			},
		},
	},
}

var FoodPromptSegmentBurnedOut = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_BURNED_OUT",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You may need to take a break and refuel your body with a balanced meal to combat burnout." +
		" You can prioritize rest and relaxation while also ensuring she is eating a variety of nutritious foods..",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be below -20 burned out",
			BoundPercentage: -20,
			BoundId:         "burned_out_ambitious",
			Type:            humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: -5,
			},
			"burned_out_ambitious": {
				CurrentPercentage: -20,
			},
			"agitated_patient": {
				CurrentPercentage: -5,
			},
			"hate_love": {
				CurrentPercentage: -5,
			},
		},
	},
}

var FoodPromptSegmentVeryBurnedOut = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER.String(),
	PromptSegmentId:    "FOOD_PROMPT_SEGMENT_VERY_BURNED_OUT",
	PromptSegmentSetId: "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
	Message: "You can recognize that taking care of her physical health is especially important when feeling very burned out. " +
		"You can reach out for support and prioritize simple, healthy meals that don't require much preparation.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER,
	RuleParts: []*humanize_protobuf.RulePart{
		{
			Description:     "must be below -35 burned out",
			BoundPercentage: -35,
			BoundId:         "burned_out_ambitious",
			Type:            humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
		},
	},
	IdealEmotionalState: &humanize_protobuf.EmotionalState{
		EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
			"sad_happy": {
				CurrentPercentage: -10,
			},
			"burned_out_ambitious": {
				CurrentPercentage: -35,
			},
			"agitated_patient": {
				CurrentPercentage: -10,
			},
			"hate_love": {
				CurrentPercentage: -10,
			},
		},
	},
}
