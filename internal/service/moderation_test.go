package service

import (
	"testing"
)

func TestCheckSensitiveWords(t *testing.T) {
	keywords := "test"
	matchedWords, count := CheckSensitiveWords(keywords)

	if count != 0 {
		t.Errorf("expected 0 matches, got %d", count)
	}

	if len(matchedWords) != 0 {
		t.Errorf("expected no matched words, got %v", matchedWords)
	}
}
