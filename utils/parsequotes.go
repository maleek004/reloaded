package utils

func ParseQuotes(word string, index int, text []string) {
	start := index
	var stop int
	for i := index + 1; i < len(text); i++ {
		if text[i] == "'" {
			stop = i
			break
		}
	}

	text[start+1] = "'" + text[start+1]
	InsertEmptyString(start, text)
	//text[start] = ""
	appendToPreviousWord(stop, text, "'")
	InsertEmptyString(stop, text)
	//text[stop] = ""
}
