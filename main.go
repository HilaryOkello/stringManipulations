package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/manipulations"
)

func main() {
	// Check if enough arguments have been provided, and if not, provide the usage
	if len(os.Args) != 3 {
		fmt.Println("Program Requires Two Arguments Passed: go run . sample.txt result.txt")
		return
	} else if os.Args[1] != "sample.txt" {
		fmt.Println("The First Argument Passed Must Be sample.txt")
		return
	} else if os.Args[2] != "result.txt" {
		fmt.Println("The Second Argument Passed Must Be result.txt")
		return
	}

	// Store the input and output files in variables
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read from the input file
	content, err := os.ReadFile(inputFile)
	if len(content) == 0 {
		fmt.Println("The sample.txt File is Empty")
		return
	}
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Split the read string (content) by space and store it as a slice of string  in variable words
	words := strings.Fields(string(content))

	// Manipulate words by passing it as an argument to the manipulateContent function and deal with errors
	manipulatedWords, err := manipulations.ManipulateContent(words)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Pass manipulatedWords to the handlePuncuation function to handle the puncuation
	punctuatedWords := manipulations.HandlePunctuation(manipulatedWords)

	// Join the final words by space and add a new line at the end
	result := strings.Join(punctuatedWords, " ") + "\n"

	// Write our result into the output file
	err = os.WriteFile(outputFile, []byte(result), 0o644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	} else {
		// Print a statement to alert us that our result has been saved in the output file
		fmt.Println("Result saved to", outputFile)
	}
}
