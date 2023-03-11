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
	narrativeTopic string,
) (*Session, error) {
	entity := &Session{
		SessionID:       sessionId,
		SpeakerName:     speakerName,
		SpeakerEntityId: speakerEntityId,
		NpcEntityIds:    strings.Join(npcEntityIds, ","),
		IsAsyncSession:  isAsync,
		WaitForCommit:   waitForCommit,
		NarrativeTopic:  narrativeTopic,
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

func (h *HumanizeDbImpl) UpdateSession(session Session, upperTx *gorm.DB) error {
	updateSession := func(tx *gorm.DB) error {
		err := tx.Updates(session)
		if err != nil {
			tx.Rollback()
			return err.Error
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = updateSession(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = updateSession(tx)
			if err != nil {
				tx.Rollback()
				return err
			}

			return nil
		})
	}
	if err != nil {
		return err
	}
	return nil
}
