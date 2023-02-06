package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(quickSort([]int{6,5,7,3,8,9,1,2,4}, 0, 8))
	// fmt.Println(quickSort([]int{6,5,7,3,8,9,1,2,4,17,13,18,19,11,12,14}, 0, 15))
	fmt.Println(binarySearch([]int{6,5,7,3,8,9,1,2,4}, 0))
	fmt.Println(binarySearch([]int{6,5,7,3,8,9,1,2,4,17,13,18,19,11,12,14}, 13))
}

func binarySearch(slc []int, val int) bool {
	slc = quickSort(slc, 0, len(slc)-1)
	i := 0
	j := len(slc) - 1

	for i < j {
		mid := int(math.Floor(float64((i+j)/2)))
		if slc[mid] == val {
			return true
		} else if slc[mid] > val {
			j = mid - 1
		} else if slc[mid] < val {
			i = mid + 1
		}
	}

	return false
}

func quickSort(slc []int, start, end int) []int {
	if start < end {
		j := partition(slc, start, end)
		quickSort(slc, start, j-1)
		quickSort(slc, j+1, end)
	}
	return slc
}

func partition(slc []int, start, end int) int {
	// Choose pivot (not optimal)
	pivot := slc[start]

	// swap pivot value with end value
	slc[start], slc[end] = slc[end], slc[start]

	j := start
	for i := start; i < end; i++ {
		if slc[i] < pivot {
			slc[i], slc[j] = slc[j], slc[i]
			j++
		}
	}
	slc[end], slc[j] = slc[j], slc[end]
	return j
}
