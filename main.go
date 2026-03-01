package main

import (
	"fmt"
	"os"
	"reloaded/utils"
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

	// Process input bytes
	finalString := utils.ParseInput(content)

	//Write processed text to output file
	err = os.WriteFile(outputFile, []byte(finalString), 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}

}
