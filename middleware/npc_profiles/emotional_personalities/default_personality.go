package emotional_personalities

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

var DefaultPersonality = []humanize_protobuf.EmotionUpdateRule{
	{
		// SLAP
		Name: "kick_rule", TriggeringAction: "SLAP", PercentageOfProc: 95,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "",
		EmotionMagnitudes: map[string]humanize_protobuf.EmotionShiftMagnitude{
			"agitated_patient": humanize_protobuf.EmotionShiftMagnitude_NEUTRAL,
			"hate_love":        humanize_protobuf.EmotionShiftMagnitude_NEUTRAL,
			"sad_happy":        humanize_protobuf.EmotionShiftMagnitude_NEUTRAL,
		},
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 20,
				BoundId:          "agitated_patient",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  30,
			},
			{
				PercentageOfProc: 20,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  30,
			},
			{
				Description:      "10% chance that we have to be at least slightly agitated",
				PercentageOfProc: 10,
				BoundId:          "agitated_patient",
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
			},
		},
	},
	{
		Name: "abuse triggered", TriggeringAction: "ABUSE_TRIGGERED", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_HIGH,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "TEXT:Please be more interesting...",
	},

	{
		Name: "admiration npc text", TriggeringAction: "NPC_EMOTION:admiration", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_HIGH,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "",
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 30,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  -36,
			},
			{
				PercentageOfProc: 30,
				BoundId:          "sad_happy",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  -36,
			},
		},
	},
	{
		Name: "amusement npc text", TriggeringAction: "NPC_EMOTION:amusement", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "",
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 30,
				BoundId:          "agitated_patient",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  -36,
			},
			{
				PercentageOfProc: 30,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  -36,
			},
		},
	},

	{
		Name: "approval npc text", TriggeringAction: "NPC_EMOTION:approval", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "",
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 20,
				BoundId:          "agitated_patient",
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
			},
		},
	},

	{
		Name: "caring npc text", TriggeringAction: "NPC_EMOTION:caring", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "",
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 100,
				BoundId:          "agitated_patient",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  -36,
			},
			{
				PercentageOfProc: 100,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  -36,
			},
			{
				PercentageOfProc: 10,
				BoundId:          "sad_happy",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  -36,
			},
		},
	},

	{
		Name: "desire npc text", TriggeringAction: "NPC_EMOTION:desire", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "",
		RuleParts:       []*humanize_protobuf.RulePart{},
	},

	{
		Name: "excitement npc text", TriggeringAction: "NPC_EMOTION:excitement", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "",
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 10,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
			},
			{
				PercentageOfProc: 10,
				BoundId:          "sad_happy",
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
			},
		},
	},

	{
		Name: "gratitude npc text", TriggeringAction: "NPC_EMOTION:gratitude", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_UNKNOWN,
		},
		UsesComposure: true, UsesLikeness: true,
		ResultingAction: "",
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 10,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
			},
			{
				PercentageOfProc: 10,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -20,
			},
			{
				PercentageOfProc: 10,
				BoundId:          "sad_happy",
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -20,
			},
		},
	},

	{
		Name: "joy npc text", TriggeringAction: "NPC_EMOTION:joy", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
				BoundId:          "hate_love",
			},
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
				BoundId:          "agitated_patient",
			},
			{
				PercentageOfProc: 10,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
			},
			{
				PercentageOfProc: 10,
				BoundId:          "sad_happy",
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
			},
		},
	},

	{
		Name: "love npc text", TriggeringAction: "NPC_EMOTION:love", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 15,
				BoundId:          "hate_love",
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
			},
		},
	},

	{
		Name: "optimism npc text", TriggeringAction: "NPC_EMOTION:optimism", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_POSITIVE,
				BoundId:          "hate_love",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -45,
				BoundId:          "hate_love",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -45,
				BoundId:          "sad_happy",
			},
		},
	},
	{
		Name: "pride npc text", TriggeringAction: "NPC_EMOTION:pride", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -40,
				BoundId:          "burned_out_ambitious",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -45,
				BoundId:          "agitated_patient",
			},
		},
	},
	{
		Name: "relief npc text", TriggeringAction: "NPC_EMOTION:relief", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -40,
				BoundId:          "burned_out_ambitious",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -30,
				BoundId:          "agitated_patient",
			},
		},
	},
	{
		Name: "anger npc text", TriggeringAction: "NPC_EMOTION:anger", PercentageOfProc: 80,
		ResultingAction: "TRIGGER_AGGRESSION",
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 50,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "sad_happy",
			},
			{
				PercentageOfProc: 50,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "burned_out_ambitious",
			},
		},
	},
	{
		Name: "anger npc text", TriggeringAction: "NPC_EMOTION:anger", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
		},
		EmotionMagnitudes: map[string]humanize_protobuf.EmotionShiftMagnitude{
			"agitated_patient": humanize_protobuf.EmotionShiftMagnitude_NEUTRAL,
			"sad_happy":        humanize_protobuf.EmotionShiftMagnitude_NEUTRAL,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "hate_love",
			},
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "burned_out_ambitious",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
				BoundId:          "sad_happy",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -40,
				BoundId:          "burned_out_ambitious",
			},
		},
	},
	{
		Name: "annoyance npc text", TriggeringAction: "NPC_EMOTION:annoyance", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "sad_happy",
			},
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "burned_out_ambitious",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -35,
				BoundId:          "burned_out_ambitious",
			},
		},
	},
	{
		Name: "disappointment npc text", TriggeringAction: "NPC_EMOTION:disappointment", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "sad_happy",
			},
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "burned_out_ambitious",
			},
		},
	},
	{
		Name: "disapproval npc text", TriggeringAction: "NPC_EMOTION:disapproval", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "burned_out_ambitious",
			},
		},
	},
	{
		Name: "embarrassment npc text", TriggeringAction: "NPC_EMOTION:embarrassment", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_UNKNOWN,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "sad_happy",
			},
			{
				PercentageOfProc: 10,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "burned_out_ambitious",
			},
		},
	},
	{
		Name: "fear npc text", TriggeringAction: "NPC_EMOTION:fear", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "sad_happy",
			},
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "burned_out_ambitious",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
				BoundId:          "agitated_patient",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
				BoundId:          "burned_out_ambitious",
			},
		},
	},
	{
		Name: "grief npc text", TriggeringAction: "NPC_EMOTION:grief", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "hate_love",
			},
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "agitated_patient",
			},
		},
	},
	{
		Name: "nervousness npc text", TriggeringAction: "NPC_EMOTION:nervousness", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
		},
		EmotionMagnitudes: map[string]humanize_protobuf.EmotionShiftMagnitude{
			"agitated_patient": humanize_protobuf.EmotionShiftMagnitude_SHORT_FUSE,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "hate_love",
			},
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "agitated_patient",
			},
		},
	},
	{
		Name: "remorse npc text", TriggeringAction: "NPC_EMOTION:remorse", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "hate_love",
			},
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "agitated_patient",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
				BoundId:          "agitated_patient",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
				BoundId:          "burned_out_ambitious",
			},
		},
	},
	{
		Name: "sadness npc text", TriggeringAction: "NPC_EMOTION:sadness", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_MEDIUM,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
		},
		EmotionMagnitudes: map[string]humanize_protobuf.EmotionShiftMagnitude{
			"sad_happy": humanize_protobuf.EmotionShiftMagnitude_NEUTRAL,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "sad_happy",
			},
			{
				PercentageOfProc: 15,
				Type:             humanize_protobuf.RulePart_IS_NEGATIVE,
				BoundId:          "burned_out_ambitious",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
				BoundId:          "sad_happy",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
				BoundId:          "hate_love",
			},
		},
	},
	{
		Name: "confusion npc text", TriggeringAction: "NPC_EMOTION:confusion", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_NEGATIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_UNKNOWN,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  10,
				BoundId:          "burned_out_ambitious",
			},
		},
	},

	{
		Name: "curiosity npc text", TriggeringAction: "NPC_EMOTION:curiosity", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -40,
				BoundId:          "agitated_patient",
			},
		},
	},

	{
		Name: "realization npc text", TriggeringAction: "NPC_EMOTION:realization", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_BELOW_X_PERCENT,
				BoundPercentage:  40,
				BoundId:          "hate_love",
			},
		},
	},
	{
		Name: "surprise npc text", TriggeringAction: "NPC_EMOTION:surprise", PercentageOfProc: 100,
		BoundShifts: map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType{
			"agitated_patient":     humanize_protobuf.EmotionUpdateRule_POSITIVE_LOW,
			"sad_happy":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"hate_love":            humanize_protobuf.EmotionUpdateRule_UNKNOWN,
			"burned_out_ambitious": humanize_protobuf.EmotionUpdateRule_UNKNOWN,
		},
		UsesComposure: true, UsesLikeness: true,
		RuleParts: []*humanize_protobuf.RulePart{
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -40,
				BoundId:          "burned_out_ambitious",
			},
			{
				PercentageOfProc: 100,
				Type:             humanize_protobuf.RulePart_IS_ABOVE_X_PERCENT,
				BoundPercentage:  -40,
				BoundId:          "sad_happy",
			},
		},
	},
}
