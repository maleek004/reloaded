package main

import (
	"fmt"
	"os"
	"reloaded/utils"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("please provide input and output file names    .... like below:")
		fmt.Println("      --> go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fmt.Println("Input: ", inputFile)
	fmt.Println("Output: ", outputFile)

	// read input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := strings.Fields(string(content))

	for index, word := range text {
		switch word {
		case "(hex)":
			res, _ := strconv.ParseInt(text[index-1], 16, 32)
			text[index-1] = strconv.Itoa(int(res))
			text[index] = ""
		case "(bin)":
			res, _ := strconv.ParseInt(text[index-1], 2, 32)
			text[index-1] = strconv.Itoa(int(res))
			text[index] = ""

		case "(up)":
			res := strings.ToUpper(text[index-1])
			text[index-1] = res
			text[index] = ""
		case "(low)":
			res := strings.ToLower(text[index-1])
			text[index-1] = res
			text[index] = ""

		case "(cap)":
			res := utils.Capitalize(text[index-1])
			text[index-1] = res
			text[index] = ""

		case ",", ".", "?", "!", ":", ";", "...":

			utils.Parsepunct(word, index, text)

		case "'":
			if index == len(text) {
				text[index-1] += "'"
				text[index] = ""
			} else if index == 0 {
				text[index+1] = "'" + text[index+1]
				text[index] = ""
			}
			if index >= 2 {
				if text[index-2] == "'" {
					text[index-1] += "'"
					text[index] = ""

				} else if text[index+2] == "'" {
					text[index+1] = "'" + text[index+1]
					text[index] = ""
				}
			} else if index <= len(text)-2 {
				if text[index+2] == "'" {
					text[index+1] += "'"
					text[index] = ""
				}

			}
		}
	}

	final := strings.Join(strings.Fields(strings.Join(text, " ")), " ")

	// Process input string

	//Write processed text to output file
	err = os.WriteFile(outputFile, []byte(final), 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}

}
