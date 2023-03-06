package prompt_management

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	ValidatorPort = "VALIDATION_PORT"
)

func ValidatePromptText(input string) (string, error) {
	validatorPort := os.Getenv(ValidatorPort)
	input = url.QueryEscape(input)
	resp, err := http.Get(fmt.Sprintf(
		"http://127.0.0.1:%s/validate?text=%s",
		validatorPort, input,
	))
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	sb := string(body)
	return sb, nil
}

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// RemoveTrailingCharacter TODO in an ideal world this should not be called
func RemoveTrailingCharacter(input string, character string, maxAllowed int) string {
	buffer := ""
	duplicateCount := 0
	processedCharacter := ""
	for _, c := range input {
		if string(c) == character {
			if processedCharacter == string(c) {
				duplicateCount++
			} else {
				duplicateCount = 0
			}
		} else {
			duplicateCount = 0
		}
		if duplicateCount < maxAllowed {
			buffer += string(c)
		}
		processedCharacter = string(c)
	}
	return buffer
}
