package prompt_management

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testserver/npc_profiles/emotional_states"
)

func TestRemoveTrailingCharacter(t *testing.T) {
	testString := "      hello"
	shouldBeString := "  hello"
	shouldBe := assert.New(t)
	output := RemoveTrailingCharacter(testString, " ", 2)
	shouldBe.Equal(shouldBeString, output)
}

func TestGetDefaultEmotionalState(t *testing.T) {
	emotionalState := emotional_states.DefaultEmotionalState
	GetEmotionalSynonym(emotionalState)
}

//if questionData == nil {
//err = g.memoryManager.CreateMessage(
//g.session.SessionID,
//effectedNpc.Entity.AskerName,
//effectedNpc.Entity.Name,
//response.Message,
//"",
//response.NpcTextEmotion,
//"",
//"",
//)
//} else {
//err = g.memoryManager.CreateMessage(
//g.session.SessionID,
//effectedNpc.Entity.AskerName,
//effectedNpc.Entity.Name,
//response.Message,
//questionData.Response,
//questionData.Emotion,
//questionData.ResponderId,
//questionData.ResponderName,
//)
