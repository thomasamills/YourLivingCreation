package db

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreateEntity(
	id, name, sessionId,
	promptId,
	generationConfigId string,
	personalityIds, promptSetIds,
	needsIds, actuationSetIds,
	religionIds, ideologyIds,
	personalityTypeIds,
	emomtionalPrimerIds []string,
) bool {
	entity := &Entity{
		EntityId:            id,
		Name:                name,
		SessionId:           sessionId,
		GenerationConfigId:  generationConfigId,
		PromptId:            promptId,
		PersonalityIds:      strings.Join(personalityIds, ","),
		PromptSetIds:        strings.Join(promptSetIds, ","),
		NeedsIds:            strings.Join(needsIds, ","),
		ActuationRuleSetIds: strings.Join(actuationSetIds, ","),
		ReligionIds:         strings.Join(religionIds, ","),
		IdeologyIds:         strings.Join(ideologyIds, ","),
		PersonalityTypeIds:  strings.Join(personalityTypeIds, ","),
	}
	_ = h.mainDB.Create(entity)
	return true
}

func (h *HumanizeDbImpl) GetEntity(entityId string) (*Entity, error) {
	result := &Entity{}
	err := h.mainDB.
		Where("entity_id = ?", entityId).
		First(&result).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println(err.Error())
		}
		return nil, err
	}

	return result, nil
}
