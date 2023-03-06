package conversation_code

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"testserver/db"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type MemoryManager interface {
	GetConversationMemoryLog(
		sessionId string,
		limitedSize int,
		currentMessage string,
		abuseFiltering bool,
		sizeLimiting bool,
	) (*humanize_protobuf.MemoryLog, error)
	CreateMessage(
		sessionId, askerName, askerId,
		responderName, responderId,
		message, response, emotion,
		responseEmotion string,
	) error
}

type MemoryManagerImpl struct {
	mu sync.Mutex
	db db.HumanizeDB
}

func (m *MemoryManagerImpl) CreateMessage(
	sessionId, askerName, askerId, responderName,
	responderId, message, response, emotion,
	responseEmotion string,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.db.CreateMessage(
		sessionId, askerName, askerId, responderName,
		responderId, message, response, emotion,
		responseEmotion,
	)
}

func NewMemoryManager(db db.HumanizeDB) MemoryManager {
	return &MemoryManagerImpl{
		db: db,
	}
}

func (m *MemoryManagerImpl) aiAbuseMechanism(
	message string,
	memLog *humanize_protobuf.MemoryLog,
) (*humanize_protobuf.MemoryLog, error) {
	// memory log processing
	sameMessageCount := 0
	if memLog != nil {
		if len(memLog.Log) > 1 {
			previousMessage := memLog.Log[len(memLog.Log)-2]
			for i := len(memLog.Log) - 2; i >= 0; i-- {
				if previousMessage.Message == message {
					sameMessageCount++
					if sameMessageCount > 1 {
						logrus.Info(fmt.Sprintf("" +
							"Ai abuse detected, " +
							"removing this message from memory",
						))
						err := m.db.DeleteMessage(memLog.Log[i].ConversationEntryId)
						if err != nil {
							return nil, err
						}
					}
				} else {
					break
				}
			}
		}
	}
	return memLog, nil
}

func reverse(s []*humanize_protobuf.ConversationEntry) []*humanize_protobuf.ConversationEntry {
	a := make([]*humanize_protobuf.ConversationEntry, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func (m *MemoryManagerImpl) limitMemoryLogSize(
	size int,
	memLog *humanize_protobuf.MemoryLog,
) (*humanize_protobuf.MemoryLog, error) {
	// memory log limiting
	newMemLog := &humanize_protobuf.MemoryLog{
		Log: make([]*humanize_protobuf.ConversationEntry, 0),
	}
	if memLog != nil {
		if len(memLog.Log) == 0 {
			return newMemLog, nil
		}
		for i := len(memLog.Log) - 1; i >= 0 && len(newMemLog.Log) <= size; i-- {
			newMemLog.Log = append(newMemLog.Log, memLog.Log[i])
		}
	}
	if len(newMemLog.Log) > 1 {
		newMemLog.Log = reverse(newMemLog.Log)
	}
	return newMemLog, nil
}

func (m *MemoryManagerImpl) GetConversationMemoryLog(
	sessionId string,
	limitedSize int,
	currentMessage string,
	abuseFiltering bool,
	sizeLimiting bool,
) (*humanize_protobuf.MemoryLog, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	memLog, err := m.db.GetMemoryLog(sessionId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"session_id": sessionId,
		}).Error(fmt.Sprintf("Error obtaining memory log"))
		return nil, err
	}
	// Perform ai abuse mechanism
	if abuseFiltering {
		memLog, err = m.aiAbuseMechanism(currentMessage, memLog)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"session_id": sessionId,
			}).Error(fmt.Sprintf("Error performing ai abuse on the memory log"))
			return nil, err
		}
	}
	// Limit memory log
	if sizeLimiting {
		memLog, err = m.limitMemoryLogSize(limitedSize, memLog)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"session_id": sessionId,
			}).Error(fmt.Sprintf("Error limiting the size of the memory log"))
			return nil, err
		}
	}
	return memLog, nil
}
