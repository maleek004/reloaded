package main

import (
	"fmt"
	"os"
	"reloaded/utils"
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
			utils.Hexa(index, text)
		case "(bin)":
			utils.Binary(index, text)

		case "(up)":
			utils.Upper(index, text)

		case "(low)":
			utils.Lower(index, text)

		case "(cap)":
			utils.Capitalize(index, text)

		case ",", ".", "?", "!", ":", ";", "...":
			utils.Parsepunct(word, index, text)

		case "'":
			utils.Parsequotes(index, text)
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
