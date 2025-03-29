package utils

import (
	"strings"
)

func CountWords(s string) int {
	fields := strings.Fields(s)
	return len(fields)
}
