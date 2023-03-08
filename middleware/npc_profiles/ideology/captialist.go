package ideology

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

var CapitalistPromptSegmentNormal = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_NORMAL",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "You are a capitalist",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentHappy = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_HAPPY",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message: "Capitalism can create opportunities for individuals and businesses to succeed" +
		" and thrive, leading to a sense of achievement and satisfaction. Its emphasis on individual " +
		"responsibility and freedom can provide a sense of empowerment and self-determination.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentVeryHappy = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_VERY_HAPPY",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message: "The success and prosperity that can be achieved under capitalism can create a" +
		" sense of optimism and possibility, inspiring people to work hard and achieve their goals.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentSad = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_SAD",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "Capitalism can be a harsh and unforgiving system that can leave many people feeling alienated and disconnected from the world around them. The focus on competition and profit can create a sense of isolation and a lack of community, leading to feelings of sadness and despair.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentVerySad = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_VERY_SAD",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "The unequal distribution of wealth and resources under capitalism can lead to widespread poverty, hunger, and suffering, leaving many people feeling hopeless and powerless to effect change.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentPatient = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_PATIENT",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "Capitalism can be an effective economic system that incentivizes innovation, competition, and efficiency. While it may have some flaws, it has been shown to create wealth and raise living standards for many people over time.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentVeryPatient = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_VERY_PATIENT",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "Despite its flaws, capitalism has been the dominant economic system in the world for centuries and has shown remarkable resilience and adaptability. Its ability to evolve and change over time has allowed it to continue to be a powerful force in the global economy.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentAgitated = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_AGITATED",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "Capitalism is an economic system that often favors the wealthy and leaves many people struggling to make ends meet. Its emphasis on profit and competition can create an unequal and unjust society where some people have much more than they need while others struggle to survive.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentVeryAgitated = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_VERY_AGITATED",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "The relentless pursuit of profit in capitalism can lead to exploitation and greed, where individuals and corporations prioritize their own financial gain over the well-being of others. This can lead to extreme income inequality, environmental degradation, and a disregard for basic human rights.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentInLove = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_IN_LOVE",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "Capitalism can be a dynamic and exciting system that provides opportunities for growth and progress. Its emphasis on innovation and competition can create a sense of excitement and possibility that many people find attractive.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentMuchInLove = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_MUCH_IN_LOVE",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "The flexibility and adaptability of capitalism can be incredibly appealing to those who are passionate about entrepreneurship and innovation. The ability to create and build new businesses and industries can be a source of great satisfaction and fulfillment.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentHatrid = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_HATE",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "The inherent inequalities and injustices of capitalism can inspire feelings of anger and resentment in those who feel left behind by the system. The focus on profit and competition can create a sense of ruthless self-interest that many people find repugnant.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentMuchHatrid = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_MUCH_HATE",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "The extreme disparities of wealth and power that can be found under capitalism can create a sense of rage and fury in those who feel oppressed or exploited by the system. The emphasis on individualism and self-interest can lead to a sense of isolation and bitterness that can be difficult to overcome.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentAmbitious = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_AMBITIOUS",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "Capitalism rewards hard work and innovation, making it a powerful motivator for those with ambitious goals and aspirations. Its emphasis on competition and success can inspire individuals to strive for excellence and achieve great things.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentVeryAmbitious = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_VERY_AMBITIOUS",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message: "The potential rewards of success under capitalism can be enormous," +
		" inspiring individuals to pursue their dreams and achieve greatness." +
		" The possibility of achieving wealth and status can create a powerful drive for those who are extremely ambitious.",
	Type: humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentBurnedOut = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_BURNED_OUT",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "The relentless pursuit of profit and success under capitalism can be exhausting and draining, leading to feelings of burnout and disillusionment. The pressure to constantly innovate and compete can create a sense of never-ending stress and anxiety.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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

var CapitalistPromptSegmentVeryBurnedOut = &humanize_protobuf.PromptSegment{
	PercentageOfProc:   100,
	PrimerType:         humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER.String(),
	PromptSegmentId:    "CAPTITALIST_PROMPT_SEGMENT_VERY_BURNED_OUT",
	PromptSegmentSetId: "CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET",
	Message:            "The constant pressure to perform and succeed under capitalism can create a culture of overwork and exhaustion, leaving many people feeling completely burned out and unable to continue.",
	Type:               humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER,
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
