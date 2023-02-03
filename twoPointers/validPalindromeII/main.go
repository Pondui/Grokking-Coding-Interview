package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPalindromeII("madame"))
	fmt.Println(validPalindromeII("raceacar"))
	fmt.Println(validPalindromeII("dead"))
	fmt.Println(validPalindromeII("abca"))
	fmt.Println(validPalindromeII("tebbem"))
	fmt.Println(validPalindromeII("eeccccbebaeeabebccceea"))
	fmt.Println(validPalindromeII("ognfjhgbjhzkqhzadmgqbwqsktzqwjexqvzjsopolnmvnymbbzoofzbbmynvmnloposjzvqxejwqztksqwbqgmdazhqkzhjbghjfno"))
}

func validPalindromeII(str string) bool {
	// Set i, j at opposite ends
	// Move i and j in each cycle, check str[i] and str[j] equal
	// If not, check next character in
	// If equal continue, set try to false
	// If not equal a second time exit
	// Assuming ASCII again

	try := true
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			if str[i+1] == str[j] && try {
				i++
				try = false
			} else if str[i] == str[j-1] && try {
				j--
				try = false
			} else {
				return false
			}
		}
	}
	return true
}
