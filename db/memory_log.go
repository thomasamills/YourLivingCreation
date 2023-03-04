package db

import (
	"fmt"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) GetMemoryLog(
	sessionId string,
) (*humanize_protobuf.MemoryLog, error) {
	var messages []MessageEntry
	result := h.mainDB.
		Where("session_id = ?", sessionId).
		Find(&messages).
		Order("message_entry_id asc")
	if result.Error != nil {
		if len(messages) == 0 {
			if result.Error == gorm.ErrRecordNotFound {
				fmt.Println(result.Error.Error())
			}
			return nil, result.Error
		}
	}
	if len(messages) == 0 {
		return nil, nil
	}
	memoryLog := &humanize_protobuf.MemoryLog{}
	for _, message := range messages {
		convEntry := &humanize_protobuf.ConversationEntry{
			AskerName:           message.AskerName,
			AskerId:             message.AskerId,
			ConversationEntryId: message.MessageEntryId,
			Message:             message.Message,
			Response:            message.Response,
			ResponderName:       message.ResponderName,
			ResponderId:         message.ResponderId,
			MessageEmotion:      message.Emotion,
			ResponseEmotion:     message.ResponseEmotion,
		}
		memoryLog.Log = append(memoryLog.Log, convEntry)
	}
	return memoryLog, nil
}
