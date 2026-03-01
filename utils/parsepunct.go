package utils

import "slices"

func Parsepunct(s string, index int, text []string) {
	punctuations := []rune{'.', ',', '?', '!', ':', ';'}
	runes := []rune(s)

	if len(runes) == 0 || !slices.Contains(punctuations, runes[0]) || index == 0 {
		return
	}

	// handling cases where we have just one punctuation
	if len(runes) == 1 {
		appendToPreviousWord(index, text, string(runes[0]))
		text[index] = ""
		return
	}

	// handling other cases: where we have just multiple punctuations like ..., or multiple punctuations merged into the next word, like !Hello

	for i, char := range runes {
		if slices.Contains(punctuations, char) || char == 39 {
			appendToPreviousWord(index, text, string(char)) // Adding te punctuation to the previous word in the slice of texts
		} else {
			// as soon as we find a character that is not a punctuation, update our slice of words with the remaining charcater from the first non punctuation, and return the function
			text[index] = string(runes[i:])
			return
		}

	}
	text[index] = ""
}
