package db

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreateSession(
	sessionId,
	speakerName,
	speakerEntityId string,
	npcEntityIds []string,
	isAsync bool,
	waitForCommit bool,
	startNarrative bool,
) (*Session, error) {
	entity := &Session{
		SessionID:       sessionId,
		SpeakerName:     speakerName,
		SpeakerEntityId: speakerEntityId,
		NpcEntityIds:    strings.Join(npcEntityIds, ","),
		IsAsyncSession:  isAsync,
		WaitForCommit:   waitForCommit,
		StartNarrative:  startNarrative,
	}
	result := h.mainDB.Create(entity)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return entity, nil
}

func (h *HumanizeDbImpl) GetSession(sessionId string) (*Session, error) {
	result := &Session{}
	err := h.mainDB.
		Where("session_id = ?", sessionId).
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

func (h *HumanizeDbImpl) UpdateSession(session Session) error {
	result := h.mainDB.Updates(session)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}
