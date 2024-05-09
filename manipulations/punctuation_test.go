package manipulations

import (
	"reflect"
	"strings"
	"testing"
)

func TestHandlePunctuation(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Place Punctuations Correctly",
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			name:     "Format groups of punctuations correctly",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "Place single quotes correctly",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:     "Place single quotes with multiple words correctly",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Fields(tc.input)
			result := strings.Join(HandlePunctuation(input), " ")
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected:%v\nFound:%v\n", tc.expected, result)
			}
		})
	}
}
