package main

import "fmt"

func main() {
	fmt.Println(reverseWordsInString("I like bananas"))
	fmt.Println(reverseWordsInString("We love Go"))
	fmt.Println(reverseWordsInString("To be or not to be"))
	fmt.Println(reverseWordsInString("You are amazing"))
	fmt.Println(reverseWordsInString("Hello     World"))
	fmt.Println(reverseWordsInString("Hey"))
}

// Works but not in place
func reverseWordsInString(slc string) string {
	// Iterate input slice from back
	// Use i and j to find words
	// Append words + space to new array
	// Only ASCII characters so 1 byte = 1 character, no need to use runes
	spc := []byte(" ")
	var ret string
	for i, j := len(slc)-1, len(slc)-1; i >= 0 && j >= 0; {
		if slc[i] == spc[0] && slc[j] == spc[0] {
			i--
			j--
		} else if slc[i] != spc[0] && slc[j] != spc[0] {
			if i-1 < 0 {
				ret = ret + slc[i:j+1] + " "
			}
			i--
		} else if slc[i] == spc[0] && slc[j] != spc[0] {
			ret = ret + slc[i+1:j+1] + " "
			j = i
		}
	}
	// Remove trailing space
	if ret[len(ret)-1:] == " " {
		ret = ret[:len(ret)-1]
	}
	return ret
}

// Alternative method
func reverseWordsInStringV2(str string) string {
	// Convert string to []byte since strings are immutable
	// Reverse whole string
	// Reverse every word in string
	// Remove duplicate spaces

	return ""
}
