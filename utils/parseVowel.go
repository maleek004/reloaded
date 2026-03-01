package utils

import "unicode"

func ParseVowel(index int, text []string) {
	if index == len(text)-1 {
		return
	}

	nextWord := []rune(text[index+1])
	char := unicode.ToLower(nextWord[0])

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		text[index] += "n"
	}
}
