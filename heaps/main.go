package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1,2,3,4,5,6,7,8,9}
	b := []int{5,6,-5,3,13,8,15,-2}
	buildMaxHeap(a)
	buildMaxHeap(b)
	fmt.Println(a)
	fmt.Println(b)
}

// By convention (for ease) the heap starts at index 1 instead of 0, reflected in the array
// Takes a node in the tree and bubbles it down until it reaches the appropriate level
func maxHeapify(a []int, heapSize int, index int) {
	l := 2 * index + 1
	r := 2 * index + 2
	largest := index

	// Check heap contains l then check if left is greater than root
	if l < heapSize && a[l] > a[index] {
		largest = l
	}

	if r < heapSize && a[r] > a[largest] {
		largest = r
	}

	// Swap nodes as necessary
	if largest != index {
		a[largest], a[index] = a[index], a[largest]
		// the index "largest" is no longer the biggest at this point
		maxHeapify(a, heapSize, largest)
	}
}

// Now we need to apply maxHeapify to the array
// To do this we'll start at the nodes above leaves and then move up the layers of the tree
// This ensures that the bottom-most subheaps (l, r, n) are already heapified, so when we iterate our tree
// Going back up (from right to left ensuring that child nodes are heapified before their parents)
// We can assume that the subheaps of the node we are on are already heaps
// And we just need to check if we must bubble down the current node
// The number of non tree nodes in a heap is given by ceil(heapsize/2)
func buildMaxHeap(a []int) {
	heapSize := len(a)
	idx := int(math.Ceil(float64(heapSize/2 - 1)))

	for i := idx; i >= 0; i-- {
		maxHeapify(a, heapSize, i)
	}
}
