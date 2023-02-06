package main

import "fmt"

func main() {
	fmt.Println(minimumWindowSubsequence("XAYMBAZABDECEAFBEABE", "ABE"))
	fmt.Println(minimumWindowSubsequence("abcdebdde", "bde"))
	fmt.Println(minimumWindowSubsequence("fgrqsqsnodwmxzkzxwqegkndaa", "kzed"))
	fmt.Println(minimumWindowSubsequence("michmznaitnjdnjkdsnmichmznait", "michmznait"))
	fmt.Println(minimumWindowSubsequence("afgegrwgwga", "aa"))
	fmt.Println(minimumWindowSubsequence("abababa", "ba"))
}

// Find shortest continuous subsequence in str1 where:
// - every character of str1 is in the substring
// - the order of the characters in substring is the same as in str2
func minimumWindowSubsequence(str1, str2 string) string {
	var (
	 	bestEnd int
	 	bestStart int
		end int
		start int
	)

	str1ptr := 0
	str2ptr := 0
	forward := true
	// Using Constraint str1.length <= 2000, to set an initial value that will always
	// be overriden on line 52, even on the first iteration
	bestEnd = 2001

	for str1ptr < len(str1) {
		// Traverse forward until you get to the end of str2
		if forward {
			if str1[str1ptr] == str2[str2ptr] {
				if str2ptr == len(str2) - 1 {
					end = str1ptr
					forward = false
				} else {
					str1ptr++
					str2ptr++
				}
			} else {
				str1ptr++
			}
		}
		// Traverse back once you have found a subsequence to find the optimal subsequence
		if !forward {
			if str1[str1ptr] == str2[str2ptr] {
				if str2ptr == 0 {
					start = str1ptr
					// Conditionally update best subsequence index
					if end - start < bestEnd - bestStart - 1 {
						bestStart = start
						bestEnd = end + 1
					}
					// Restart the algorithm from the unexplored subslice
					str1ptr = end + 1
					str2ptr = 0
					forward = true
				} else {
					str1ptr--
					str2ptr--
				}
			} else {
				str1ptr--
			}
		}
	}
	// Confirmation to check the leftmost minimum-length window is actually returned
	// fmt.Println("Start index: ", bestStart)
	// fmt.Println("End index: ", bestEnd - 1)
	return str1[bestStart: bestEnd]
}

// Brute force approach is to start with window of length str2
// for every possible substring in str1, check if it meets the condition
// if not, increase window size and continue
// repeat until window size is equal to the length of str1
