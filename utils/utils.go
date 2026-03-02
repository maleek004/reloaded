package utils

import (
	"strconv"
	"strings"
)

func Capitalize(index int, text []string) {
	runes := []rune(text[index-1])
	runes[0] = runes[0] - 32
	text[index-1] = string(runes)
	text[index] = ""
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

func Lower(index int, text []string) {
	res := strings.ToLower(text[index-1])
	text[index-1] = res
	text[index] = ""
}

func Upper(index int, text []string) {
	res := strings.ToUpper(text[index-1])
	text[index-1] = res
	text[index] = ""
}

func Binary(index int, text []string) {
	res, _ := strconv.ParseInt(text[index-1], 2, 32)
	text[index-1] = strconv.Itoa(int(res))
	text[index] = ""
}

func Hexa(index int, text []string) {
	res, _ := strconv.ParseInt(text[index-1], 16, 32)
	text[index-1] = strconv.Itoa(int(res))
	text[index] = ""
}

func Parsequotes(index int, text []string) {
	if (index == len(text)) || (index == len(text)-1) {
		text[index-1] += "'"
		text[index] = ""
	} else if index <= 1 {
		text[index+1] = "'" + text[index+1]
		text[index] = ""
	} else {
		if strings.HasSuffix(text[index+1], "'") || text[index+2] == "'" {
			text[index+1] = "'" + text[index+1]
			text[index] = ""
		} else if strings.HasPrefix(text[index-1], "'") || text[index-2] == "'" {
			text[index-1] += "'"
			text[index] = ""
		}
	}
}
