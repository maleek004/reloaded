package utils

func Capitalize(s string) string {
	runes := []rune(s)
	runes[0] = runes[0] - 32

	return string(runes)
}

func Parsepunct(s string, index int, text []string) {
	if text[index-2] == "" && text[index-1] == "" {
		text[index-3] += s
	} else if text[index-1] == "" {
		text[index-2] += s
	} else {
		text[index-1] += s
	}
	text[index] = ""
}
