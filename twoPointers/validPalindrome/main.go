package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("kayak"))
	fmt.Println(validPalindrome("hello"))
	fmt.Println(validPalindrome("RACEACAR"))
	fmt.Println(validPalindrome("A"))
	fmt.Println(validPalindrome("ABCDABCD"))
	fmt.Println(validPalindrome("DCBAABCD"))
	fmt.Println(validPalindrome("ABCBA"))
}

func validPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
