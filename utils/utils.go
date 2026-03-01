package utils

func InsertEmptyString(index int, text []string) {
	text[index] = ""
}

func appendToPreviousWord(index int, text []string, s string) {
	//Skips Possible Empty Spaces in our slice of words while trying to append string to the previous word
	for i := index - 1; i >= 0; i-- {
		if text[i] != "" {
			text[i] += s
			break
		}
	}
}

// func S2PES(index int, text []string, s string) { //Skips 2 Possible Empty Spaces in our slice of words while trying to append string to the previous word
// 	if text[index-2] == "" && text[index-1] == "" {
// 		text[index-3] += s
// 	} else if text[index-1] == "" {
// 		text[index-2] += s
// 	} else {
// 		text[index-1] += s
// 	}
// }
