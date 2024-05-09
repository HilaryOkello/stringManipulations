package manipulations

import (
	"reflect"
	"strings"
	"testing"
)

func TestManipulateContent(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Replace hexadecimal numbers",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "Replace binary numbers",
			input:    "10 (bin) years",
			expected: "2 years",
		},
		{
			name:     "Convert word to uppercase",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO !",
		},
		{
			name:     "Convert word to lowercase",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			name:     "Convert word to capitalized",
			input:    "Welcome to the Brooklyn bridge (cap)",
			expected: "Welcome to the Brooklyn Bridge",
		},
		{
			name:     "Convert multiple words to uppercase",
			input:    "This is so exciting (up, 2)",
			expected: "This is SO EXCITING",
		},
		{
			name:     "Convert 'a' to 'an' before vowels",
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Fields(tc.input)
			result := strings.Join(ManipulateContent(input), " ")

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v\nFound: %v\n", tc.expected, result)
			}
		})
	}
}
