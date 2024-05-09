package manipulations

import (
	"strings"
)

func HandlePunctuation(words []string) []string {
	punctuationSet := map[string]bool{".": true, ",": true, "!": true, "?": true, ":": true, "'": true, ";": true}
	for i := 0; i < len(words); i++ {
		word := words[i]

		if punctuationSet[string(word[0])] {
			if word == "'" {
				// Find the index of the closing single quote
				closingQuoteIndex := findClosingQuoteIndex(words[i+1:])
				if closingQuoteIndex != -1 {
					// append the next word at index [i+1] to the opening quote
					words[i] += words[i+1]
					// Remove that next word
					words = append(words[:i+1], words[i+2:]...)
					// concatenate the closing quote to the previous word
					words[i+closingQuoteIndex-1] += words[i+closingQuoteIndex]
					// Remove the closing quote
					words = append(words[:i+closingQuoteIndex], words[i+1+closingQuoteIndex:]...)
				}
			} else {
				isAllPunct := true
				for j := 0; j < len(word); j++ {
					if punctuationSet[string(word[j])] {
						// Append punctuation to the previous word
						if len(words[i-1]) > 0 {
							words[i-1] += string(word[j])
						}
					} else {
						// Update the current word if not punctuation
						if j > 0 {
							words[i] = word[j:]
						}
						isAllPunct = false
						break
					}
				}
				// If the whole word is punctuation, remove it
				if isAllPunct {
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
		}
	}
	return words
}

func findClosingQuoteIndex(words []string) int {
	for i, word := range words {
		if strings.HasSuffix(word, "'") {
			return i
		}
	}
	return -1
}
