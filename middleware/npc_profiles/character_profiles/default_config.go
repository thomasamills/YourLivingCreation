package character_profiles

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

var DefaultConfig = &humanize_protobuf.GenerationConfig{
	GenerationConfigId:     "DEFAULT_CONFIG",
	Temperature:            0.6,
	TopP:                   1.0,
	TopK:                   60,
	RepetitionPenalty:      1.1,
	WordsInEmotionalPrimer: 5,
	Length:                 200,
	StopSequences: []string{
		"\n",
		"\n\n",
	},
	MaxMemoryLogEntries: 3,
}
