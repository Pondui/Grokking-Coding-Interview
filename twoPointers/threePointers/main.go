package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{1,-1,0}, -1))
	fmt.Println(threeSum([]int{3,7,1,2,8,4,5}, 10))
	fmt.Println(threeSum([]int{3,7,1,2,8,4,5}, 20))
	fmt.Println(threeSum([]int{3,7,1,2,8,4,5}, 21))
	fmt.Println(threeSum([]int{-1,2,1,-4,5,-3}, -8))
	fmt.Println(threeSum([]int{-1,2,1,-4,5,-3}, 0))
	fmt.Println(threeSum([]int{-1,2,1,-4,5,-3}, 7))
}

func threeSum(slc []int, target int) bool {
	// Sort slice
	// Fix i element
	// Instantiate j to lower, k to upper
	// Reassign j and k as necessary
	// Repeat

	sort.Ints(slc)
	for i := 0; i < len(slc)-2; i++ {
		for j, k := i+1, len(slc)-1; j < k; {
			if slc[i] + slc[j] + slc[k] < target {
				j++
			} else if slc[i] + slc[j] + slc[k] > target {
				k--
			} else {
				return true
			}
		}
	}
	return false
}
