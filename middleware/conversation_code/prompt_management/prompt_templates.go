package prompt_management

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"testserver/db"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
	"text/template"
	"time"
)

type PromptTemplate struct {
	AskerName             *string
	ResponderName         *string
	EmotionalPrimer       *string
	MotivationalPrimer    *string
	IdeologyPrimer        *string
	ReligionPrimer        *string
	PersonalityTypePrimer *string
	MemLog                *string
	Synonym               *string
	Topic                 *string
	StoryBackground       *string
}

func ParseStructIntoTemplate(input interface{}, templ string) ([]byte, error) {
	temp, err := template.New("").Option("missingkey=zero").Parse(templ)
	if err != nil {
		fmt.Println(templ)
		return nil, err
	}
	templateBuffer := new(bytes.Buffer)
	err = temp.Execute(templateBuffer, input)
	if err != nil {
		fmt.Println(templ)
		return nil, err
	}
	return templateBuffer.Bytes(), nil
}

func GenerateMemoryBuffer(
	memLog *humanize_protobuf.MemoryLog,
	askerName, responderName string,
) (string, error) {
	memoryLogString := ""
	for _, conversationEntry := range memLog.Log {
		if len(conversationEntry.Response) == 0 {
			memoryLogString += fmt.Sprintf(
				"$$%s: %s",
				conversationEntry.AskerName, conversationEntry.Message,
			)
		} else {
			memoryLogString += fmt.Sprintf(
				"$$%s: %s$%s: %s",
				conversationEntry.AskerName, conversationEntry.Message,
				conversationEntry.ResponderName, conversationEntry.Response,
			)
		}
	}
	return memoryLogString, nil
}

func GeneratePromptFromList(
	askerName, responderName, promptTemplate string,
	promptSegmentToAdd []*humanize_protobuf.PromptSegment,
	memLog *humanize_protobuf.MemoryLog,
	emotionalState *humanize_protobuf.EmotionalState,
	session db.Session,
) (string, error) {
	currentPrompt := promptTemplate
	// Then loop through every single prompt segment in the array omitting ones of the same type
	segmentsToAdd := &PromptTemplate{
		AskerName:     &askerName,
		ResponderName: &responderName,
	}
	if memLog != nil {
		memoryLogString, err := GenerateMemoryBuffer(memLog, askerName, responderName)
		if err != nil {
			return "", err
		}
		segmentsToAdd.MemLog = &memoryLogString
		memoryLogStringBytes, err := ParseStructIntoTemplate(segmentsToAdd, memoryLogString)
		if err != nil {
			return "", err
		}
		memoryLogString = string(memoryLogStringBytes)
	}
	for _, promptSegment := range promptSegmentToAdd {
		switch promptSegment.Type {
		case humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_EMOTIONAL_PRIMER:
			// First get the current emotional synonym, this uses the {{.Synonym}} placeholder
			synonym, err := GetEmotionalSynonym(emotionalState)
			if err != nil {
				return "", err
			}
			segmentsToAdd.Synonym = &synonym
			emotionalPrimer, err := ParseValuesIntoPrimer(segmentsToAdd, promptSegment.Message)
			if err != nil {
				return "", err
			}
			segmentsToAdd.EmotionalPrimer = &emotionalPrimer
			break
		case humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_STORY_BACKGROUND:
			storyBackgroundPrimer, err := ParseValuesIntoPrimer(segmentsToAdd, promptSegment.Message)
			if err != nil {
				return "", err
			}
			segmentsToAdd.StoryBackground = &storyBackgroundPrimer
			break
		case humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_MOTIVATIONAL_PRIMER:
			motivationalPrimer, err := ParseValuesIntoPrimer(segmentsToAdd, promptSegment.Message)
			if err != nil {
				return "", err
			}
			segmentsToAdd.MotivationalPrimer = &motivationalPrimer
			break
		case humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_IDEOLOGY_PRIMER:
			ideologyPrimer, err := ParseValuesIntoPrimer(segmentsToAdd, promptSegment.Message)
			if err != nil {
				return "", err
			}
			segmentsToAdd.IdeologyPrimer = &ideologyPrimer
			break
		case humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_RELIGION_PRIMER:
			religionPrimer, err := ParseValuesIntoPrimer(segmentsToAdd, promptSegment.Message)
			if err != nil {
				return "", err
			}
			segmentsToAdd.ReligionPrimer = &religionPrimer
			break
		case humanize_protobuf.PromptSegmentType_PROMPT_SEGMENT_TYPE_PERSONALITY_TYPE_PRIMER:
			personalityTypePrimer, err := ParseValuesIntoPrimer(segmentsToAdd, promptSegment.Message)
			if err != nil {
				return "", err
			}
			segmentsToAdd.PersonalityTypePrimer = &personalityTypePrimer
			break
		}
	}
	if len(session.NarrativeTopic) > 0 {
		str := "The main topic of conversation which you will always refer to is: " + session.NarrativeTopic
		segmentsToAdd.Topic = &str
	}
	byteArray, err := ParseStructIntoTemplate(segmentsToAdd, currentPrompt)
	if err != nil {
		return "", err
	}
	text := StandardizeSpaces(string(byteArray))
	return text, nil
}

func ParseValuesIntoPrimer(promptTemplateStruct *PromptTemplate, emotionalPrimer string) (string, error) {
	byteArray, err := ParseStructIntoTemplate(promptTemplateStruct, emotionalPrimer)
	if err != nil {
		return "", err
	}
	return string(byteArray), nil
}

func GetEmotionalSynonym(
	emotionalState *humanize_protobuf.EmotionalState,
) (string, error) {
	// To store the selected synonyms
	selectedSynonyms := make([]string, 0)
	for _, bound := range emotionalState.EmotionalBounds {
		roundedCurrentPercentage :=
			math.Round(float64(bound.CurrentPercentage/10)) * 10
		if synonymsForEmotion, ok := bound.Synonyms[int32(roundedCurrentPercentage)]; ok {
			selectedSynonyms = append(selectedSynonyms, synonymsForEmotion.Synonyms[0])
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(selectedSynonyms), func(i, j int) {
		selectedSynonyms[i],
			selectedSynonyms[j] = selectedSynonyms[j],
			selectedSynonyms[i]
	})
	return selectedSynonyms[0], nil
}
