package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Capitalize(word string) string {
	if word == "" {
		return ""
	}
	return strings.Title(word)
}

func GetNPreviousWords(s string) int {
	num, err := strconv.Atoi(strings.TrimSuffix(s, ")"))
	if err != nil {
		fmt.Println("HelperFunctionError --> extracting number from '", s, "' failed 🚧, defaulting to 1")
		return 1
	}
	return num
}

func ApplyToPrevWords(index int, text []string, n int, action func(string) string, multiple bool) {
	for i := index - 1; i >= 0 && n > 0; i-- {
		if text[i] == "" {
			continue
		}

		text[i] = action(text[i])
		n--

	}

	InsertEmptyString(index, text)
	if multiple {
		InsertEmptyString(index+1, text)
	}
}
