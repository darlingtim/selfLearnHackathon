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
		inputContent, err := os.ReadFile(outputFile)
		// check and capture reading input file error
		if err != nil {
			return fmt.Errorf("unable to read file: %v", err), ""
		}

		inputContentString := string(inputContent)

	}
}
