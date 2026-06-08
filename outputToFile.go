package main

import (
	"fmt"
	"os"
	"strings"
)

func outputToFile(inputFile string, outputFile string) (error, string) {
	// check for wrong file format.
	if !(strings.HasSuffix(inputFile, ".txt") && strings.HasSuffix(outputFile, ".txt")) {
		return fmt.Errorf("please enter the right file formats: input and output file must be a text file which ends with '.txt'"), ""

	} else {
		// read input file
		inputContent, err := os.ReadFile(inputFile)
		// check and capture reading input file error
		if err != nil {
			return fmt.Errorf("unable to read file: %v", err), ""
		}

		// check if input file is empty
		if len(inputContent) == 0 {
			return fmt.Errorf("input file is empty. Please, ensure there is a content in the file, and try again."), ""
		}

		inputContentString := string(inputContent)

		// capitalise every file text, line by line
		var result strings.Builder
		for _, line := range strings.Split(inputContentString, "\n") {
			result.WriteString(strings.ToUpper(line) + "\n")
		}

		// write to output file
		writeError := os.WriteFile(outputFile, []byte(result.String()), 0644)
		if writeError != nil {
			return writeError, ""
		}

	}
	return nil, fmt.Sprintf("%v has been successfully read, Capitalized and written to %v", inputFile, outputFile)
}
