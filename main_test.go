package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Store out test cases in a []struct
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Test Case 1",
			input:          "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.\n",
			expectedOutput: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.\n",
		},
		{
			name:           "Test case 2",
			input:          "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.\n",
			expectedOutput: "Simply add 66 and 2 and you will see the result is 68.\n",
		},
		{
			name:           "Test case 3",
			input:          "There is no greater agony than bearing a untold story inside you.\n",
			expectedOutput: "There is no greater agony than bearing an untold story inside you.\n",
		},
		{
			name:           "Test case 4",
			input:          "Punctuation tests are ... kinda boring ,don't you think !?\n",
			expectedOutput: "Punctuation tests are... kinda boring, don't you think!?\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Write test input to a temporary file
			inputFile := "sample.txt"
			err := os.WriteFile(inputFile, []byte(tc.input), 0o644)
			if err != nil {
				t.Fatal("Error creating test input file:", err)
			}

			// Prepare command to execute the main program
			testCommand := exec.Command("go", "run", ".", inputFile, "result.txt")
			// Declare variabled to capture the standard output & error of our testCommand
			var stdout, stderr bytes.Buffer
			testCommand.Stdout = &stdout
			testCommand.Stderr = &stderr

			// Execute the command
			err = testCommand.Run()
			if err != nil {
				t.Fatalf("Error running the program: %s, %s\n", err, stderr.String())
			}
			// Read output filet
			output, err := os.ReadFile("result.txt")
			if err != nil {
				t.Fatalf("Error reading output file: %s\n", err)
			}

			// Check if output matches the expected output
			if reflect.DeepEqual(output, tc.expectedOutput) {
				t.Errorf("Output mismatch for %q\nExpected: %q\nFound: %q", tc.name, tc.expectedOutput, output)
			}
		})
	}
}
