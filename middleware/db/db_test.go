package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

func TestHumanizeDbImpl_test_encoding_of_synonyms(t *testing.T) {
	shouldBe := assert.New(t)
	sut := &HumanizeDbImpl{}
	testSynonyms := map[int32]*humanize_protobuf.SynonymList{
		-50: {
			Synonyms: []string{
				"Extremely Sad", "Extremely Unhappy", "Extremely Miserable",
				"Extremely Sorrowful", "Extremely Dejected",
				"Extremely Woeful", "Extremely Low-Spirited",
				"Extremely Down", "Extremely Dismal",
			},
		},
		-40: {
			Synonyms: []string{
				"Very Sad", "Very Unhappy", "Very Miserable",
				"Very Sorrowful", "Very Dejected",
				"Very Woeful", "Very Low-Spirited",
				"Very Down", "Very Dismal",
			},
		},
	}
	encodedSynonyms := sut.encodeSynonymMap(testSynonyms)
	decodedSynonyms, err := sut.decodeSynonymMap(encodedSynonyms)
	if err != nil {
		t.Fail()
	}
	for key, synonymList := range testSynonyms {
		for index, synonym := range synonymList.Synonyms {
			shouldBe.Equal(synonym, decodedSynonyms[key].Synonyms[index])
		}
	}
}

func TestHumanizeDbImpl_test_encoding_of_bound_shifts(t *testing.T) {
	shouldBe := assert.New(t)
	sut := &HumanizeDbImpl{}
	testBoundShifts := map[string]humanize_protobuf.EmotionShiftMagnitude{
		"sad_happy":         30,
		"patience_agitated": 30,
	}
	encodedSynonyms := sut.encodeStringInt32Map(testBoundShifts)
	decodedSynonyms, err := sut.decodeStringInt32Map(encodedSynonyms)
	if err != nil {
		t.Fail()
	}
	shouldBe.Equal(testBoundShifts, decodedSynonyms)
}

func TestHumanizeDbImpl_test_encoding_of_actions(t *testing.T) {
	shouldBe := assert.New(t)
	sut := &HumanizeDbImpl{}
	testBoundShifts := map[string]humanize_protobuf.NodeTraversalAction{
		"path1": humanize_protobuf.NodeTraversalAction_AGITATES_NPC,
		"path2": humanize_protobuf.NodeTraversalAction_SENTIMENT_ACTION,
	}
	encodedSynonyms := sut.encodeActions(testBoundShifts)
	decodedSynonyms := sut.decodeActions(encodedSynonyms)

	shouldBe.Equal(testBoundShifts, decodedSynonyms)
}
