package manipulations

import (
	"errors"
	"strconv"
	"strings"
)

func ManipulateContent(words []string) ([]string, error) {
	// We loop through words and perform several operations depending with the case that matches
	for i := 0; i < len(words); i++ {
		word := words[i]

		// if strings.Contains(word, "cap,") {
		// 	num, err := strconv.Atoi(strings.Trim(word, "(cap,)"))
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	fmt.Println(num)
		// }
		switch word {

		case "(bin)":
			if i > 0 {
				// convert the previous word before (bin)
				val, err := strconv.ParseInt(words[i-1], 2, 64)
				if err != nil {
					return nil, errors.New("Error converting binary to decimal: " + err.Error())
				}
				words[i-1] = strconv.Itoa(int(val))
			}
			words = append(words[:i], words[i+1:]...)

		case "(hex)":
			if i > 0 {
				val, err := strconv.ParseInt(words[i-1], 16, 64)
				if err != nil {
					return nil, errors.New("Error converting hexadecimal to decimal: " + err.Error())
				}
				words[i-1] = strconv.Itoa(int(val))
			}
			words = append(words[:i], words[i+1:]...)

		case "(cap)":
			if i > 0 {
				words[i-1] = strings.ToUpper(string(words[i-1][0])) + string(words[i-1][1:])
			}
			words = append(words[:i], words[i+1:]...)

		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)

		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}

			words = append(words[:i], words[i+1:]...)

		case "(cap,":
			numStr := string(words[i+1][:len(words[i+1])-1])
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, errors.New("character passed to (cap, <number>) not a number: " + err.Error())
			}
			if len(words[:i]) >= num {
				for j := 1; j <= num; j++ {
					words[i-j] = strings.ToUpper(string(words[i-j][0])) + string(words[i-j][1:])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				return nil, errors.New("there isn't enough words to manipulate based on the <number> stipulated")
			}

		case "(up,":
			numStr := string(words[i+1][:len(words[i+1])-1])
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, errors.New("Character passed to (up, <number>) not a number: " + err.Error())
			}
			if len(words[:i]) >= num {
				for j := 1; j <= num; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				return nil, errors.New("there isn't enough words to manipulate based on the <number> stipulated")
			}

		case "(low,":
			numStr := string(words[i+1][:len(words[i+1])-1])
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, errors.New("Character passed to (low, <number>) not a number: " + err.Error())
			}
			if len(words[:i]) >= num {
				for j := 1; j <= num; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
			} else {
				return nil, errors.New("there isn't enough words to manipulate based on the <number> stipulated")
			}

		case "a":
			vowels := "aeiouAEIOUhH"
			for _, vowel := range vowels {
				if i < len(words)-1 && strings.HasPrefix(words[i+1], string(vowel)) {
					words[i] = "an"
				}
			}
		case "A":
			vowels := "aeiouAEIOUhH"
			for _, vowel := range vowels {
				if i < len(words)-1 && strings.HasPrefix(words[i+1], string(vowel)) {
					words[i] = "An"
				}
			}
		}
	}

	return words, nil
}
