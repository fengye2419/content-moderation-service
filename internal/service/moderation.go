package service

import (
	"strings"

	"github.com/fengye2419/content-moderation-service/internal/model"
)

// CheckSensitiveWords checks the given keywords for sensitive words.
func CheckSensitiveWords(keyword string) ([]string, int) {
	sensitiveWords, _ := model.GetSensitiveWords()
	matchedWords := []string{}
	matchCount := 0
	keywordStr := string(keyword) // Convert rune to string
	for _, sensitiveWord := range sensitiveWords {
		if strings.Contains(keywordStr, sensitiveWord.Word) {
			matchedWords = append(matchedWords, sensitiveWord.Word)
			matchCount++
		}
	}

	return matchedWords, matchCount
}
