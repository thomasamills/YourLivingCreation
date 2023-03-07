package personality_types

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

var ISTJPersonalityPromptSegmentSegmentNormal = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_NORMAL",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "ISTJ (Introverted, Sensing, Feeling, Judging) - kind, compassionate, and dependable.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentHappy = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_HAPPY",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "ISTJs may not show their happiness as outwardly as some other personality types, but they can still experience a sense of satisfaction and contentment when they complete a task or achieve a goal. They may take pride in their work and feel a sense of accomplishment when they see the results of their efforts.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentVeryHappy = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_VERY_HAPPY",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message: "When an ISTJ accomplishes something significant, they may feel a" +
		" great sense of joy and happiness. They may celebrate their success and take the" +
		" time to appreciate what they have accomplished.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentSad = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_SAD",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message: "ISTJs can become sad when their plans or expectations are not met. They may have a hard time " +
		"dealing with change, especially if it was not something they had prepared for. They may withdraw and become introspective, trying to figure out what went wrong and how they can avoid similar situations in the future.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentVerySad = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_VERY_SAD",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "When faced with significant loss or disappointment, an ISTJ may struggle to cope. They may become overly critical of themselves and others, and they may struggle to see a way forward.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentPatient = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_PATIENT",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "When faced with significant loss or disappointment, an ISTJ may struggle to cope. They may become overly critical of themselves and others, and they may struggle to see a way forward.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentVeryPatient = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_VERY_PATIENT",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "When they are committed to a project or goal, ISTJs can be extremely patient, persevering through setbacks and obstacles. They may be willing to put in long hours and go above and beyond what is expected of them to ensure that things are done correctly.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentAgitated = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_AGITATED",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "The ISTJ's attention to detail can sometimes make them frustrated when others don't take things as seriously as they do. They may become angry when their plans or routines are disrupted, and they may struggle to adapt to new situations.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentVeryAgitated = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_VERY_AGITATED",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "When pushed too far, an ISTJ can become very angry and may express their frustration in a cold, logical manner. They may become critical and unyielding, and their stubbornness may cause them to clash with others.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentInLove = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_IN_LOVE",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "ISTJs may struggle to express their emotions, but when they are in love, they can be extremely loyal and committed partners. They may show their love through actions rather than words, taking care of their partner and making sure their needs are met.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentMuchInLove = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_MUCH_IN_LOVE",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            " When an ISTJ is deeply in love, they may become more expressive and affectionate than usual. They may go out of their way to do things for their partner and may become more open about their feelings.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentHatrid = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_HATE",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "The ISTJ may become extremely critical and judgmental when consumed with hatred towards someone or something. They may fixate on every little detail that they dislike and may have a hard time letting go of the negative emotions. They may become closed off to new ideas and experiences, and may refuse to deviate from their set ways of doing things.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentMuchHatrid = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_MUCH_HATE",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "When consumed with extreme hatred, the ISTJ may become completely rigid in their beliefs and actions. They may become obsessed with the object of their hatred, constantly analyzing and scrutinizing every detail. They may have difficulty trusting others and may become paranoid that others are out to get them. In some cases, they may even become vengeful and seek retribution against those they hate.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentAmbitious = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_AMBITIOUS",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "ISTJs can be highly ambitious when it comes to achieving their goals. They may set high standards for themselves and work tirelessly to meet them. They may have a clear vision of what they want to accomplish and be willing to do whatever it takes to get there.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentVeryAmbitious = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_VERY_AMBITIOUS",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message: " When an ISTJ is very ambitious, they may be driven to succeed at all costs. They may be willing" +
		" to make sacrifices and take risks to achieve their goals, even if it means putting their personal relationships or well-being on the line.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentBurnedOut = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_BURNED_OUT",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "ISTJs can become burned out when they take on too many responsibilities or when they feel like they are not making progress towards their goals. They may become exhausted and disengaged, feeling like they are stuck in a rut.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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

var ISTJPersonalityPromptSegmentSegmentVeryBurnedOut = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER.String(),
	PromptSegmentId:    "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SEGMENT_VERY_BURNED_OUT",
	PromptSegmentSetId: "ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET",
	Message:            "When an ISTJ is extremely burned out, they may become withdrawn and uncommunicative. They may struggle to find the motivation to do anything and may feel like they are in a deep rut that they can't get out of.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER,
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
