package model

import (
	"testing"
)

func TestSensitiveWords(t *testing.T) {
	words, _ := GetSensitiveWords()
	if len(words) == 0 {
		t.Errorf("expected some sensitive words, got none")
	}
}
