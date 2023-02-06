package main

import "math"

func main() {

}

func maxHeapify(a []int, heapSize int, index int) {
	l := int(math.Pow(2, float64(index)))
	r := int(math.Pow(2, float64(index))) + 1

	largest := index

	// First check if l is even in the heap
	if l < heapSize && a[l] > a[index] {
		largest = l
	}

	if r < heapSize && a[r] > a[index] {
		largest = r
	}

	// Swap nodes as necessary
	if largest != index {
		a[largest], a[index] = a[index], a[largest]
		// the index largest is no longer the biggest at this point
		maxHeapify(a, heapSize, largest)
	}
}
