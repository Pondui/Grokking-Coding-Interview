package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	fmt.Println(maxInSlidingWindowDeque(3, []int{5,6,3,-3,0,2,1,4,10,-2}))
	fmt.Println(maxInSlidingWindowDeque(3, []int{1,2,3,4,5,6,7,8,9,10}))
	fmt.Println(maxInSlidingWindowDeque(4, []int{3,3,3,3,3,3,3,3,3,3}))
	fmt.Println(maxInSlidingWindowDeque(2, []int{10,6,9,-3,23,-1,34,56,67,-1,-4,-8,-2,9,10,34,67}))
	fmt.Println(maxInSlidingWindowDeque(1, []int{4,5,6,1,2,3}))
	fmt.Println(maxInSlidingWindowDeque(2, []int{9,5,3,1,6,3}))
	fmt.Println(maxInSlidingWindowDeque(2, []int{1,2}))
}

// The idea here is not avoid calculating the max element in each window
// Instead we'd like to keep track of the current max element and only update it when:
// (a) it leaves the window, or
// (b) a greater element enters the window
// We'd also like to keep elements of the window in general because once the max
// element leaves the window, another element can be promoted to the max 
// (we can't do this if we don't track all elements)
// That being said, once a greater element than it's predecessors
// (but not necessarily greater than the current max) enters the window, we can safely discard the predecessors
// since they will never be the max element in the current and future windows
type data struct {
	value int
	index int
}

func maxInSlidingWindowDeque(windowSize int, slc []int) []int {
	res := []int{}
	var dq deque.Deque[data]

	for i, val := range slc {
		if dq.Len() > 0 && i - dq.Front().index >= windowSize {
			dq.PopFront()
		}
		for dq.Len() > 0 && dq.Back().value < val {
			dq.PopBack()
		}
		dq.PushBack(data{
			value: val,
			index: i,
		})
		if i >= windowSize - 1 {
			res = append(res, dq.Front().value)
		}
	}

	return res
}

// This can also be implemented using a max heap, where we initialise the heap
// with values from the first window and then add/rm from the heap as we slide
// func maxInSlidingWindowHeap(windowSize int, slc []int) []int {
// 	res := []int{}
// 	return res
// }
