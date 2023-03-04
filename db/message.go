package db

import "fmt"

func (h *HumanizeDbImpl) CreateMessage(
	sessionId,
	askerName,
	askerId,
	responderName,
	responderId,
	message,
	response,
	emotion string,
	responseEmotion string,
) error {
	m := &MessageEntry{
		SessionId:       sessionId,
		AskerName:       askerName,
		AskerId:         askerId,
		ResponderName:   responderName,
		ResponderId:     responderId,
		Message:         message,
		Response:        response,
		Emotion:         emotion,
		ResponseEmotion: responseEmotion,
	}
	result := h.mainDB.Create(m)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

func (h *HumanizeDbImpl) DeleteMessage(messageId int32) error {
	err := h.mainDB.Where("message_entry_id = ?", messageId).Delete(&MessageEntry{})
	if err != nil {
		if err.Error != nil {
			return err.Error
		}
	}
	return nil
}
