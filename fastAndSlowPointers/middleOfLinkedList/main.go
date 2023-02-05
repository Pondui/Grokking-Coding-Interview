package main

import "fmt"

func main() {
	ll1 := linkedListFromSlice([]int{0,1})
	ll2 := linkedListFromSlice([]int{0,1,2,3,4})
	ll3 := linkedListFromSlice([]int{0,1,2,3,4,5})
	ll4 := linkedListFromSlice([]int{0,1,2,3,4,5,6,7,8,9})
	ll5 := linkedListFromSlice([]int{3,2,1})
	ll6 := linkedListFromSlice([]int{10})

	fmt.Println(middleOfLinkedList(&ll1))
	fmt.Println(middleOfLinkedList(&ll2))
	fmt.Println(middleOfLinkedList(&ll3))
	fmt.Println(middleOfLinkedList(&ll4))
	fmt.Println(middleOfLinkedList(&ll5))
	fmt.Println(middleOfLinkedList(&ll6))
}

func middleOfLinkedList(ll *LinkedList) int {
	i, j := ll.head, ll.head

	for j != nil {
		// Need to check j.next to avoid nil pointer dereference
		if j.next != nil {
			j = j.next.next
			i = i.next
		} else {
			j = nil
		}
	}

	return i.data
}
