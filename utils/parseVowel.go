package utils

import (
	"strings"
	"unicode"
)

func ParseVowel(index int, text []string) {
	if index == len(text)-1 {
		return
	}
	var nextWord []rune
	for i := index + 1; i < len(text); i++ {
		if !strings.HasPrefix(text[i], "(") && !strings.HasSuffix(text[i], ")") {
			nextWord = []rune(text[i])
			break

		}

	}

	char := unicode.ToLower(nextWord[0])

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		text[index] += "n"
	}
}
