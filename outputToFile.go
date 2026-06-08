package main

import (
	"fmt"
	"strings"
)

func outputToFile(inputFile string, outputFile string) (error, string) {
	// check for wrong file format.
	if !(strings.HasSuffix(inputFile, ".txt") && strings.HasSuffix(outputFile, ".txt")) {
		return fmt.Errorf("please enter the right file formats: input and output file must be a text file which ends with '.txt'"), ""

	}
}
