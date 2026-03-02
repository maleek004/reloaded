package utils

import (
	"slices"
	"strings"
)

func ParseInput(content string) string {
	text := strings.Fields(content)

	for index, word := range text {
		switch word {
		case "(hex)":
			text[index-1] = ParseInt(index, text, 16)

		case "(bin)":
			text[index-1] = ParseInt(index, text, 2)

		case "(up)":
			ApplyToPrevWords(index, text, 1, strings.ToUpper, false)

		case "(low)":
			ApplyToPrevWords(index, text, 1, strings.ToLower, false)

		case "(cap)":
			ApplyToPrevWords(index, text, 1, Capitalize, false)

		case "(up,":
			ApplyToPrevWords(index, text, GetNPreviousWords(text[index+1]), strings.ToUpper, true)

		case "(low,":
			ApplyToPrevWords(index, text, GetNPreviousWords(text[index+1]), strings.ToLower, true)

		case "(cap,":
			ApplyToPrevWords(index, text, GetNPreviousWords(text[index+1]), Capitalize, true)

		case "'":
			ParseQuotes(word, index, text)

		case "a", "A":
			ParseVowel(index, text)

		default:
			Parsepunct(word, index, text)

		}
	}

	text = slices.DeleteFunc(text, func(s string) bool { return s == "" })
	final := strings.Join(text, " ")

	return final
}
