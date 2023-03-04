package emotional_states

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

var DefaultEmotionalState = &humanize_protobuf.EmotionalState{
	EntityId:      "DEFAULT_EMOTIONAL_STATE",
	PersonalityId: "DEFAULT_PERSONALITY",
	Composure:     80,
	Likeness:      100,
	EmotionalBounds: map[string]*humanize_protobuf.EmotionalBound{
		"sad_happy": {
			EntityId:              "DEFAULT_EMOTIONAL_STATE",
			BoundId:               "sad_happy",
			CurrentPercentage:     0,
			UpperBound:            "happy",
			LowerBound:            "sad",
			EmotionShiftMagnitude: humanize_protobuf.EmotionShiftMagnitude_BOTTLING_UP,
			Synonyms: map[int32]*humanize_protobuf.SynonymList{
				-50: {
					Synonyms: []string{
						"Extremely Sad", "Extremely Unhappy", "Extremely Miserable",
						"Extremely Sorrowful", "Extremely Dejected",
						"Extremely Woeful", "Extremely Low-Spirited",
						"Extremely Down", "Extremely Dismal",
					},
				},
				-40: {
					Synonyms: []string{
						"Very Sad", "Very Unhappy", "Very Miserable",
						"Very Sorrowful", "Very Dejected",
						"Very Woeful", "Very Low-Spirited",
						"Very Down", "Very Dismal",
					},
				},
				-30: {
					Synonyms: []string{
						"Sad", "Unhappy", "Miserable",
						"Sorrowful", "Dejected",
						"Woeful", "Low-Spirited",
						"Down", "Dismal",
					},
				},
				-20: {
					Synonyms: []string{
						"Slightly Sad", "Slightly Unhappy", "Slightly Miserable",
						"Slightly Sorrowful", "Slightly Dejected",
						"Slightly Woeful", "Slightly Low-Spirited",
						"Slightly Down", "Slightly Dismal",
					},
				},
				-10: {
					Synonyms: []string{
						"A Little Sad", "A Little Unhappy", "A Little Miserable",
						"A Little Sorrowful", "A Little Dejected",
						"A Little Woeful", "A Little Low-Spirited",
						"A Little Down", "A Little Dismal",
					},
				},
				0: {
					Synonyms: []string{
						"Neutral",
					},
				},
				10: {
					Synonyms: []string{
						"Slightly Happy", "Slightly Satisfied", "Slightly Content",
						"Slightly Joyful", "Slightly Jovial",
					},
				},
				20: {
					Synonyms: []string{
						"Quite Happy", "Quite Satisfied", "Quite Content",
						"Quite Joyful", "Quite Jovial",
					},
				},
				30: {
					Synonyms: []string{
						"Happy", "Satisfied", "Content",
						"Joyful", "Jovial",
					},
				},
				40: {
					Synonyms: []string{
						"Very Happy", "Very Satisfied", "Very Content",
						"Very Joyful", "Very Jovial",
					},
				},
				50: {
					Synonyms: []string{
						"Extremely Happy", "Extremely Satisfied", "Extremely Content",
						"Extremely Joyful", "Extremely Jovial",
					},
				},
			},
		},
		"agitated_patient": {
			EntityId:              "DEFAULT_EMOTIONAL_STATE",
			BoundId:               "agitated_patient",
			CurrentPercentage:     0,
			UpperBound:            "patient",
			LowerBound:            "agitated",
			EmotionShiftMagnitude: humanize_protobuf.EmotionShiftMagnitude_BOTTLING_UP,
			Synonyms: map[int32]*humanize_protobuf.SynonymList{
				-50: {
					Synonyms: []string{
						"Extremely Worried", "Extremely Upset", "Extremely Anxious",
						"Extremely Perturbed", "Extremely Disquiet",
						"Extremely Distressed", "Extremely Concerned",
						"Extremely Troubled", "Extremely Alarmed",
					},
				},
				-40: {
					Synonyms: []string{
						"Very Worried", "Very Upset", "Very Anxious",
						"Very Perturbed", "Very Disquiet",
						"Very Distressed", "Very Concerned",
						"Very Troubled", "Very Alarmed",
					},
				},
				-30: {
					Synonyms: []string{
						"Worried", "Upset", "Anxious",
						"Perturbed", "Disquiet",
						"Distressed", "Concerned",
						"Troubled", "Alarmed",
					},
				},
				-20: {
					Synonyms: []string{
						"Slightly Worried", "Slightly Upset", "Slightly Anxious",
						"Slightly Perturbed", "Slightly Disquiet",
						"Slightly Distressed", "Slightly Concerned",
						"Slightly Troubled", "Slightly Alarmed",
					},
				},
				-10: {
					Synonyms: []string{
						"A Little Worried", "A Little Upset", "A Little Anxious",
						"A Little Perturbed", "A Little Disquiet",
						"A Little Distressed", "A Little Concerned",
						"A Little Troubled", "A Little Alarmed",
					},
				},
				-0: {
					Synonyms: []string{
						"Neutral", "Indifferent",
					},
				},
				10: {
					Synonyms: []string{
						"A Little Tolerant", "A Little Restraintful", "A Little Resigned",
						"A Little Stoic", "A Little Composed",
						"A Little Serene", "A Little Tranquil",
						"A Little Calm", "A Little Fortitudinous",
					},
				},
				20: {
					Synonyms: []string{
						"Slightly Tolerant", "Slightly Restraintful", "Slightly Resigned",
						"Slightly Stoic", "Slightly Composed",
						"Slightly Serene", "Slightly Tranquil",
						"Slightly Calm", "Slightly Fortitudinous",
					},
				},
				30: {
					Synonyms: []string{
						"Tolerant", "Restraintful", "Resigned",
						"Stoic", "Composed",
						"Serene", "Tranquil",
						"Calm", "Fortitudinous",
					},
				},
				40: {
					Synonyms: []string{
						"Very Tolerant", "Very Restraintful", "Very Resigned",
						"Very Stoic", "Very Composed",
						"Very Serene", "Very Tranquil",
						"Very Calm", "Very Fortitudinous",
					},
				},
				50: {
					Synonyms: []string{
						"Extremely Tolerant", "Extremely Restraintful", "Extremely Resigned",
						"Extremely Stoic", "Extremely Composed",
						"Extremely Serene", "Extremely Tranquil",
						"Extremely Calm", "Extremely Fortitudinous",
					},
				},
			},
		},
		"burned_out_ambitious": {
			EntityId:              "DEFAULT_EMOTIONAL_STATE",
			BoundId:               "burned_out_ambitious",
			CurrentPercentage:     0,
			UpperBound:            "ambitious",
			LowerBound:            "burned_out",
			EmotionShiftMagnitude: humanize_protobuf.EmotionShiftMagnitude_BOTTLING_UP,
			Synonyms: map[int32]*humanize_protobuf.SynonymList{
				-50: {
					Synonyms: []string{
						"Very Unambitious", "Very Burnt-out", "Very Drained",
						"Very Worn", "Very Jaded",
					},
				},
				-40: {
					Synonyms: []string{
						"Very Unambitious", "Very Burnt-out", "Very Drained",
						"Very Worn", "Very Jaded",
					},
				},
				-30: {
					Synonyms: []string{
						"Unambitious", "Burnt-out", "Drained",
						"Worn", "Jaded",
					},
				},
				-20: {
					Synonyms: []string{
						"Slightly Unambitious", "Slightly Burnt-out", "Slightly Drained",
						"Slightly Worn", "Slightly Jaded",
					},
				},
				-10: {
					Synonyms: []string{
						"A Little Unambitious", "A Little Burnt-out", "A Little Drained",
						"A Little Worn", "A Little Jaded",
					},
				},
				-0: {
					Synonyms: []string{
						"Neutral", "Indifferent",
					},
				},
				10: {
					Synonyms: []string{
						"A Little Eager", "A Little Ambitious", "A Little Aspiring", "A Little Motivated",
						"A Little Driven",
					},
				},
				20: {
					Synonyms: []string{
						"Slightly Ambitious", "Slightly Aspiring", "Slightly Motivated",
						"Slightly Driven", "Slightly Eager",
					},
				},
				30: {
					Synonyms: []string{
						"Ambitious", "Aspiring", "Motivated",
						"Driven", "Eager",
					},
				},
				40: {
					Synonyms: []string{
						"Very Tolerant", "Very Restraintful", "Very Resigned",
						"Very Stoic", "Very Composed",
						"Very Serene", "Very Tranquil",
						"Very Calm", "Very Fortitudinous",
					},
				},
				50: {
					Synonyms: []string{
						"Extremely Tolerant", "Extremely Restraintful", "Extremely Resigned",
						"Extremely Stoic", "Extremely Composed",
						"Extremely Serene", "Extremely Tranquil",
						"Extremely Calm", "Extremely Fortitudinous",
					},
				},
			},
		},
		"hate_love": {
			EntityId:              "DEFAULT_EMOTIONAL_STATE",
			BoundId:               "hate_love",
			CurrentPercentage:     0,
			UpperBound:            "love",
			LowerBound:            "hate",
			EmotionShiftMagnitude: humanize_protobuf.EmotionShiftMagnitude_BOTTLING_UP,
			Synonyms: map[int32]*humanize_protobuf.SynonymList{
				-50: {
					Synonyms: []string{
						"Extremely Standoffish",
						"Extremely Untrusting",
						"Extremely Sarcastic",
						"Extremely Hostile",
						"Extremely Unwelcoming",
					},
				},
				-40: {
					Synonyms: []string{
						"Very Standoffish",
						"Very Untrusting",
						"Very Sarcastic",
						"Very Hostile",
						"Very Unwelcoming",
					},
				},
				-30: {
					Synonyms: []string{
						"Standoffish",
						"Untrusting",
						"Sarcastic",
						"Hostile",
						"Unwelcoming",
					},
				},
				-20: {
					Synonyms: []string{
						"Quite Standoffish",
						"Quite Untrusting",
						"Quite Sarcastic",
						"Quite Hostile",
						"Quite Unwelcoming",
					},
				},
				-10: {
					Synonyms: []string{
						"Slightly Standoffish",
						"Slightly Untrusting",
						"Slightly Sarcastic",
						"Slightly Hostile",
						"Slightly Unwelcoming",
					},
				},
				0: {
					Synonyms: []string{
						"Neutral",
					},
				},
				10: {
					Synonyms: []string{
						"Slightly Approachable",
						"Slightly Caring",
						"Slightly Understanding",
						"Slightly Smitten",
						"Slightly Passionate",
						"Slightly Welcoming",
					},
				},
				20: {
					Synonyms: []string{
						"Quite Approachable",
						"Quite Caring",
						"Quite Understanding",
						"Quite Smitten",
						"Quite Passionate",
						"Quite Welcoming",
					},
				},
				30: {
					Synonyms: []string{
						"Approachable",
						"Caring",
						"Understanding",
						"Smitten",
						"Passionate",
						"Welcoming",
					},
				},
				40: {
					Synonyms: []string{
						"Very Approachable",
						"Very Caring",
						"Very Understanding",
						"Very Smitten",
						"Very Passionate",
						"Very Welcoming",
					},
				},
				50: {
					Synonyms: []string{
						"Extremely Approachable",
						"Extremely Caring",
						"Extremely Understanding",
						"Extremely Smitten",
						"Extremely Passionate",
						"Extremely Welcoming",
					},
				},
			},
		},
	},
	Fears:    []string{},
	Hobbies:  []string{},
	IsPreset: true,
}
