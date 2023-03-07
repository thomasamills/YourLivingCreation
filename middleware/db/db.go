package db

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testserver/db/id_gen"
	"testserver/npc_profiles/character_profiles"
	"testserver/npc_profiles/emotional_states"
	"testserver/npc_profiles/needs"
	"testserver/npc_profiles/prompts"
	"testserver/npc_profiles/rule_sets"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type HumanizeDB interface {
	Init() error

	CreateEntity(
		id, name, sessionId,
		personalityId,
		generationConfigId, promptSetId,
		promptSegmentSetId, promptId,
		actuationSetId, automationPromptSetId,
		automationPromptSetSegmentId string,
	) bool
	GetEntity(
		entityId string,
	) (*Entity, error)
	UpdateEntityNeeds(
		entityId string,
		needs []string,
	) (*Entity, error)
	CreateSession(
		sessionId,
		speakerName,
		speakerEntityId string,
		npcEntityIds []string,
		isAsync bool,
		waitForCommit bool,
		startNarrative bool,
	) (*Session, error)
	GetSession(
		sessionId string,
	) (*Session, error)
	UpdateSession(
		session Session,
	) error

	CreateMessage(
		sessionId,
		askerName,
		askerId,
		responderName,
		responderId,
		message,
		response,
		emotion,
		responseEmotion string,
	) error
	DeleteMessage(
		messageId int32,
	) error

	GetMemoryLog(
		sessionId string,
	) (*humanize_protobuf.MemoryLog, error)

	CreateGenerationConfig(
		id string,
		config *humanize_protobuf.GenerationConfig,
		upperTx *gorm.DB,
	) error
	ListConfigs() ([]string, error)
	GetGenerationConfig(
		configId string,
	) (*humanize_protobuf.GenerationConfig, error)
	DeleteGenerationConfig(
		genConfigId string,
	) error

	CreateEmotionalState(
		entityId string, isPreset bool,
		state *humanize_protobuf.EmotionalState,
		upperTx *gorm.DB,
	) error
	CopyEmotionalState(
		presetId, entityId string,
	) (*humanize_protobuf.EmotionalState, error)
	GetEmotionalState(
		stateId string,
		upperTx *gorm.DB,
	) (*humanize_protobuf.EmotionalState, error)
	DeleteEmotionalState(
		emotionalStateId string,
		upperTx *gorm.DB,
	) error

	SavePersonality(
		personalityId string,
		rules []humanize_protobuf.EmotionUpdateRule,
		upperTx *gorm.DB,
	) error
	GetPersonality(
		personalityId string,
		upperTx *gorm.DB,
	) ([]*humanize_protobuf.EmotionUpdateRule, error)
	ListPersonalities() ([]string, error)
	DeletePersonality(
		personalityId string,
		upperTx *gorm.DB,
	) error

	CreateEmotionalBound(
		bound *humanize_protobuf.EmotionalBound,
		boundId string,
		entityId string, // if we are overriding the entity id
		upperTx *gorm.DB,
	) (string, error)
	UpdateEmotionalBound(
		boundInstanceId string,
		emotionalBound *humanize_protobuf.EmotionalBound,
	) (string, error)
	DeleteEmotionalBound(
		boundInstanceId string,
		upperTx *gorm.DB,
	) error
	UpdateEmotionalBoundPercentage(
		emotionalBound *humanize_protobuf.EmotionalBound,
	) error

	CreatePrompt(
		prompt *humanize_protobuf.Prompt,
		upperTx *gorm.DB,
	) (string, error)
	GetPrompt(
		promptId string,
		upperTx *gorm.DB,
	) (*humanize_protobuf.Prompt, error)
	ListPrompts(
		promptSegmentSetId string,
		upperTx *gorm.DB,
	) ([]string, error)
	UpdatePrompt(
		prompt *humanize_protobuf.Prompt,
		upperTx *gorm.DB,
	) error
	DeletePrompt(
		promptId string,
		upperTx *gorm.DB,
	) error
	ListPromptSegmentsBySetIdAndPrimerType(
		promptSegmentSetId string,
		primerType string,
		upperTx *gorm.DB,
	) ([]string, error)
	CreatePromptSegment(
		promptSegment *humanize_protobuf.PromptSegment,
		upperTx *gorm.DB,
	) (string, error)
	GetPromptSegment(
		promptId string,
		upperTx *gorm.DB,
	) (*humanize_protobuf.PromptSegment, error)
	ListPromptSegmentsBySetId(
		promptSegmentSetId string,
		upperTx *gorm.DB,
	) ([]string, error)
	UpdatePromptSegment(
		prompt *humanize_protobuf.PromptSegment,
		upperTx *gorm.DB,
	) error
	DeletePromptSegment(
		promptId string,
		upperTx *gorm.DB,
	) error

	ListPresetStates(
		upperTx *gorm.DB,
	) ([]string, error)

	GetRulePart(
		rulePartId string,
		upperTx *gorm.DB,
	) (*humanize_protobuf.RulePart, error)
	GetRuleParts(
		ruleId string,
		upperTx *gorm.DB,
	) ([]*humanize_protobuf.RulePart, error)
	CreateRuleParts(
		associatingId string,
		ruleParts []*humanize_protobuf.RulePart,
		upperTx *gorm.DB,
	) error
	DeletePersonalityRule(
		ruleId string,
		upperTx *gorm.DB,
	) error
	DeletePersonalityRulePart(
		ruleId string,
		upperTx *gorm.DB,
	) error

	CreateEmotionalBoundRule(
		rule *humanize_protobuf.EmotionUpdateRule,
	) (string, error)
	CreateEmotionalBoundRulePart(
		rule *humanize_protobuf.RulePart,
	) (string, error)

	GetPromptSet(
		promptSetId string,
		upperTx *gorm.DB,
	) ([]*humanize_protobuf.Prompt, error)
	GetPromptIdealStatesBySetId(
		promptSetId string,
		upperTx *gorm.DB,
	) ([]IdealStateAndId, error)
	ListPromptSetIds(
		upperTx *gorm.DB,
	) ([]string, error)
	GetPromptSegmentIdealStateBySetIdAndPrimerType(
		promptSegmentSetId string,
		primerType string,
		upperTx *gorm.DB,
	) ([]IdealStateAndId, error)
	GetPromptSegmentSet(
		promptSegmentSetId string,
		primerType string,
		upperTx *gorm.DB,
	) ([]*humanize_protobuf.PromptSegment, error)
	ListPromptSegmentSets(
		upperTx *gorm.DB,
	) ([]string, error)
	GetActuationRule(
		id string, upperTx *gorm.DB,
		byPersonalityId bool,
	) (*humanize_protobuf.ActuationRule, error)
	GetActuationRuleSet(
		actuationRuleSetId string,
		upperTx *gorm.DB,
	) ([]*humanize_protobuf.ActuationRule, error)
	ListActuationRules(
		personalityId string,
		upperTx *gorm.DB,
	) ([]string, error)
	CreateActuationRuleSet(
		personalityId string,
		actuationRuleSet []*humanize_protobuf.ActuationRule,
	) error
	CreateOrUpdateActuationRule(
		actuationRulePartId string,
		rule *humanize_protobuf.ActuationRule,
		upperLevelTx *gorm.DB,
		personalityId string,
	) error
	DeleteActuationRule(actuationRuleId string) error
}

type HumanizeDbImpl struct {
	mainDB *gorm.DB
	idGen  id_gen.ULIDGenerator
}

func NewHumanizeDb() HumanizeDB {
	db := &HumanizeDbImpl{
		idGen: id_gen.NewULIDGenerator(),
	}
	if err := db.Init(); err != nil {
		fmt.Println(err.Error())
	}
	// Default human
	err := db.SavePersonality(
		"DEFAULT_PERSONALITY",
		rule_sets.DefaultPersonality,
		nil,
	)
	if err != nil {
		logrus.Error("could not save the default personality")
	}
	err = db.CreateEmotionalState(
		"DEFAULT_EMOTIONAL_STATE",
		true,
		emotional_states.DefaultEmotionalState,
		nil,
	)
	if err != nil {
		logrus.Error("could not save the default emotional state")
	}
	err = db.CreateGenerationConfig(
		"DEFAULT_CONFIG",
		character_profiles.DefaultConfig,
		nil,
	)
	if err != nil {
		logrus.Error("could not save the default config")
	}

	_, err = db.CreatePrompt(prompts.FarRightPubManNormal, nil)
	if err != nil {
		logrus.Error("could not far right pub man normal prompt")
	}
	_, err = db.CreatePrompt(prompts.YoungFarLeftSoberManNormal, nil)
	if err != nil {
		logrus.Error("could not young far left sober man prompt" + err.Error())
	}
	_, err = db.CreatePromptSegment(needs.FoodPromptSegmentNormal, nil)
	if err != nil {
		logrus.Error("could food need prompt normal")
	}
	return db
}

func (h *HumanizeDbImpl) Init() error {
	db, err := gorm.Open(sqlite.Open("humanize.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	// Migrate the schema
	err = db.AutoMigrate(&Session{})
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&Entity{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&GenerationConfig{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&MessageEntry{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&EmotionalState{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&EmotionalBound{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&EmotionalBoundRule{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&RulePart{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&CurrentQuestState{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Personality{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Quest{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Prompt{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&PromptSegment{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&ActuationRule{}); err != nil {
		return err
	}
	h.mainDB = db
	return nil
}

func (h *HumanizeDbImpl) validatePromptRequest(prompt *humanize_protobuf.Prompt) (bool, string) {
	if prompt.IdealEmotionalState == nil {
		return false, "no ideal emotional state assigned to the prompt"
	}
	return true, ""
}
func (h *HumanizeDbImpl) encodePromptSegmentTypes(
	segmentTypes []humanize_protobuf.PromptSegmentType,
) string {
	segmentTypeString := ""
	for i, segmentType := range segmentTypes {
		if i != len(segmentTypes)-1 {
			segmentTypeString += strconv.Itoa(int(segmentType)) + ","
		} else {
			segmentTypeString += strconv.Itoa(int(segmentType))
		}
	}
	return segmentTypeString
}

func (h *HumanizeDbImpl) decodePromptSegmentTypes(
	encodedString string,
) ([]humanize_protobuf.PromptSegmentType, error) {
	splitStr := strings.Split(encodedString, ",")
	segmentTypes := make([]humanize_protobuf.PromptSegmentType, 0)
	if len(encodedString) < 1 {
		return segmentTypes, nil
	}
	for _, segmentType := range splitStr {
		parsedSegmentType, err := strconv.Atoi(segmentType)
		if err != nil {
			return nil, err
		}
		segmentTypes = append(segmentTypes, humanize_protobuf.PromptSegmentType(parsedSegmentType))
	}
	return segmentTypes, nil
}

func (h *HumanizeDbImpl) encodeSynonymMap(synonyms map[int32]*humanize_protobuf.SynonymList) string {
	encodedString := ""
	count := 0
	for key, value := range synonyms {
		encodedString += strconv.Itoa(int(key))
		for _, synonym := range value.Synonyms {
			encodedString += fmt.Sprintf("#%s", synonym)
		}
		if count != len(synonyms) {
			encodedString += ","
		}
	}
	return encodedString
}

func (h *HumanizeDbImpl) decodeSynonymMap(
	encodedMap string,
) (map[int32]*humanize_protobuf.SynonymList, error) {
	decodedMap := make(map[int32]*humanize_protobuf.SynonymList, 0)
	encodedLists := strings.Split(encodedMap, ",")
	for _, synonymList := range encodedLists {
		decodedList := strings.Split(synonymList, "#")
		if len(decodedList) > 1 {
			decodedSynonymMapKey, err := strconv.Atoi(decodedList[0])
			if err != nil {
				return nil, errors.New("could not parse map key from encoded synonym string")
			}
			decodedMap[int32(decodedSynonymMapKey)] = &humanize_protobuf.SynonymList{
				Synonyms: decodedList[1:],
			}
		}
	}
	return decodedMap, nil
}

func (h *HumanizeDbImpl) encodeStringInt32Map(
	boundShifts map[string]humanize_protobuf.EmotionShiftMagnitude,
) string {
	if boundShifts == nil {
		return ""
	}
	encodedBoundShiftString := ""
	count := 1
	for key, value := range boundShifts {
		encodedBoundShiftString += fmt.Sprintf("%s:%d", key, int(value))
		if count != len(boundShifts) {
			encodedBoundShiftString += ","
		}
		count++
	}
	return encodedBoundShiftString
}

func (h *HumanizeDbImpl) decodeStringInt32Map(encodedBoundShifts string,
) (map[string]humanize_protobuf.EmotionShiftMagnitude, error) {
	if encodedBoundShifts == "" {
		return nil, nil
	}
	decodedBoundShift := make(map[string]humanize_protobuf.EmotionShiftMagnitude, 0)
	boundShifts := strings.Split(encodedBoundShifts, ",")
	for _, boundShift := range boundShifts {
		splitBoundShift := strings.Split(boundShift, ":")
		boundId := splitBoundShift[0]
		boundShiftChange, err := strconv.Atoi(splitBoundShift[1])
		if err != nil {
			return nil, errors.New("could not parse bound shift percentage")
		}
		decodedBoundShift[boundId] = humanize_protobuf.EmotionShiftMagnitude(boundShiftChange)
	}
	return decodedBoundShift, nil
}

func (h *HumanizeDbImpl) encodeBoundShiftMap(
	boundShifts map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType,
) string {
	if boundShifts == nil {
		return ""
	}
	encodedBoundShiftString := ""
	count := 1
	for key, value := range boundShifts {
		encodedBoundShiftString += fmt.Sprintf("%s:%d", key, int(value))
		if count != len(boundShifts) {
			encodedBoundShiftString += ","
		}
		count++
	}
	return encodedBoundShiftString
}

func (h *HumanizeDbImpl) decodeBoundShiftMap(encodedBoundShifts string,
) (map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType, error) {
	if encodedBoundShifts == "" {
		return nil, nil
	}
	decodedBoundShift := make(map[string]humanize_protobuf.EmotionUpdateRule_BoundShiftType, 0)
	boundShifts := strings.Split(encodedBoundShifts, ",")
	for _, boundShift := range boundShifts {
		splitBoundShift := strings.Split(boundShift, ":")
		boundId := splitBoundShift[0]
		boundShiftChange, err := strconv.Atoi(splitBoundShift[1])
		if err != nil {
			return nil, errors.New("could not parse bound shift percentage")
		}
		decodedBoundShift[boundId] = humanize_protobuf.EmotionUpdateRule_BoundShiftType(boundShiftChange)
	}
	return decodedBoundShift, nil
}
