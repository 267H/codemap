package utils

import (
	"strings"
)

func EstimateTokens(text string) int {
	charCount := len(text)

	whitespaceCount := 0
	newlineCount := 0
	symbolCount := 0

	for _, char := range text {
		if strings.ContainsRune(" \t\r", char) {
			whitespaceCount++
		} else if char == '\n' {
			newlineCount++
		} else if !isAlphaNumeric(char) {
			symbolCount++
		}
	}

	alphaNumCount := charCount - whitespaceCount - newlineCount - symbolCount

	wordTokens := int(float64(alphaNumCount) * 0.17)
	whitespaceTokens := whitespaceCount / 4
	newlineTokens := newlineCount
	symbolTokens := symbolCount

	return wordTokens + whitespaceTokens + newlineTokens + symbolTokens
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_'
}

func CountWords(s string) int {
	fields := strings.Fields(s)
	return len(fields)
}
