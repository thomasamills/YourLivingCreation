package npcdata

import (
	"testserver/db"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
	"time"
)

type NpcData struct {
	EmotionalState       *humanize_protobuf.EmotionalState
	GptConfig            *humanize_protobuf.GenerationConfig
	ActuationRuleSet     []*humanize_protobuf.ActuationRule
	Personality          []*humanize_protobuf.EmotionUpdateRule
	LastInputTime        time.Time
	AverageInputInterval time.Duration
	NpcRequestChannel    chan *ActionRequest
	NpcStopChannel       chan bool
	Entity               *db.Entity
	IsPaused             bool
}

type ActionRequest struct {
	PythonRequest *humanize_protobuf.HumanizeRequest
	EffectedNpc   *NpcData
	ReturnResult  bool
	ReturnChannel chan *humanize_protobuf.HumanizeResponse
}
