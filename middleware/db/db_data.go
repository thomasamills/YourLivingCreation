package db

type EmotionalState struct {
	EntityId                string `gorm:"primaryKey;autoIncrement:false;unique"`
	PersonalityID           string
	Composure               int32
	Likeness                int32
	Fears                   string // in csv format
	Hobbies                 string // in csv format
	IsPreset                bool   // if it is a preset it will be immutables
	EmotionalSynonymCounter int32
}

type Personality struct {
	PersonalityId string `gorm:"primaryKey;autoIncrement:false;unique"`
}

type EmotionalBound struct {
	BoundInstanceId         string `gorm:"primaryKey;autoIncrement:false;unique"`
	EntityId                string
	BoundId                 string
	CurrentPercentage       int64
	LowerBound              string
	UpperBound              string
	EncodedSynonymMap       string
	EmotionalShiftMagnitude int32
}

type EmotionalBoundRule struct {
	RuleID                       string `gorm:"primaryKey;autoIncrement:false;unique"`
	PersonalityID                string
	Name                         string
	PercentageOfProc             int32
	TriggeringAction             string
	ResultingAction              string
	BoundShifts                  string // csv map in format '<bound-id>:<percentage_change>'
	UsesComposure                bool
	UsesLikeness                 bool
	EncodedShiftMagnitudeChanges string // csv map in format '<bound-id>:<shift-magnitude-enum>'
}

type RulePart struct {
	RulePartId               string `gorm:"primaryKey;autoIncrement:false;unique"`
	RuleID                   string
	PercentageOfProc         int32
	RuleType                 int32 // corresponds to protobuf enum
	BoundPercentage          int32
	BoundId                  string
	ConversationEmotionTrain string // csv
}

type Session struct {
	SessionID       string `gorm:"primaryKey;autoIncrement:false;unique"`
	SpeakerName     string
	SpeakerEntityId string
	NpcEntityIds    string // csv
	IsAsyncSession  bool
	WaitForCommit   bool
	StartNarrative  bool
}

type Entity struct {
	EntityId                              string `gorm:"primaryKey;autoIncrement:false;unique"`
	SessionId                             string
	Name                                  string
	PersonalityId                         string
	GenerationConfigId                    string
	PromptSetId                           string
	PromptSegmentSetId                    string
	PromptId                              string
	ActuationRuleSetId                    string
	AutonomousPromptSetId                 string
	AutonomousPromptSegmentSetId          string
	KnowledgeGraphCurrentSatisfactoryNeed string
}

type MessageEntry struct {
	MessageEntryId  int32  `gorm:"primaryKey;autoIncrement:true;unique"`
	SessionId       string //fk
	AskerName       string
	AskerId         string
	ResponderName   string
	ResponderId     string
	Message         string
	Response        string
	Emotion         string
	ResponseEmotion string
}

type GenerationConfig struct {
	GenerationConfigId  string `gorm:"primaryKey;autoIncrement:false;unique"`
	Entity1Id           string
	Entity2Id           string
	Temperature         float32
	TopP                float32
	TopK                int64
	RepetitionPenalty   float32
	Length              int64
	StopSequences       string
	CharacterPrompt     string
	NumberOfSynonyms    int32
	MaxMemoryLogEntries int32
}

type CurrentQuestState struct {
	// key path to the next possible node
	QuestInstanceId string `gorm:"primaryKey;autoIncrement:false;unique"`
	QuestId         string
	QuestPath       string
	EncodedActions  string
}

type QuestTreeNode struct {
	QuestTreeNodeId      string `gorm:"primaryKey;autoIncrement:false;unique"`
	ParentNodeQuestPath  string
	QuestPath            string
	QuestId              string
	ActionToArriveAtNode int32
}

type Quest struct {
	QuestId                     string `gorm:"primaryKey;autoIncrement:false;unique"`
	EntityNames                 string
	EntityPersonalities         string
	EntityGenerationConfigs     string
	EntityPresetEmotionalStates string
	RootNodePath                string
}

type Prompt struct {
	PromptId             string `gorm:"primaryKey;autoIncrement:false;unique"`
	PromptText           string
	PromptType           int32
	PromptName           string
	PromptSetId          string
	RequiredSegmentTypes string // encoded from array of proto enums
}

type PromptSegment struct {
	PromptSegmentId       string `gorm:"primaryKey;autoIncrement:false;unique"`
	Message               string
	PromptSegmentSetId    string
	PromptSegmentType     int32
	IdealEmotionalStateId string
	PercentageOfProc      int32
	PrimerType            string
}

type ActuationRule struct {
	ActuationRuleId          string `gorm:"primaryKey;autoIncrement:false;unique"`
	RuleType                 int32
	PersonalityId            string
	ResultingPromptId        string
	ResultingPromptSegmentId string
	IdealStateId             string
	PercentageOfProc         int32
	RuleName                 string
}
