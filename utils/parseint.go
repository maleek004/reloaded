package utils

import "strconv"

func ParseInt(index int, text []string, base int) string {
	res, _ := strconv.ParseInt(text[index-1], base, 32)
	InsertEmptyString(index, text)
	return strconv.Itoa(int(res))

}
