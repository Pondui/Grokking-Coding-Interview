package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(happyNumber(23))
	fmt.Println(happyNumber(2))
	fmt.Println(happyNumber(1))
	fmt.Println(happyNumber(5))
	fmt.Println(happyNumber(19))
	fmt.Println(happyNumber(25))
	fmt.Println(happyNumber(7))
}

func happyNumber(n int) bool {
	// Reduces to cycle detection in generated values n
	// Assumes n either terminates in 1 or cycles infinitely
	// Not proven why n can't continue indefinitely without cycles
	
	m := n
	for m != 1 {
		digitsN := extractDigits(n)
		k := 0
		for _, e := range digitsN {
			k += int(math.Pow(float64(e), 2))
		}
		n = k

		for j := 0; j < 2; j++ {
			digitsM := extractDigits(m)
			l := 0
			for _, e := range digitsM {
				l += int(math.Pow(float64(e), 2))
			}
			m = l
		}

		if n == m && n != 1 {
			return false
		}
	}
	return true
}

func extractDigits(n int) []int {
	var res []int
	m := float64(n)
	for m > 0 {
		res = append(res, int(m) % 10)
		m = math.Floor(m/10)
	}
	return res
}
