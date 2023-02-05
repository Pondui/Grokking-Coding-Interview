package main

import "fmt"

func main() {
	ll1 := linkedListFromSlice([]int{0,1,2,3,4,5,6,7,8,9})
	ll2 := linkedListFromSlice([]int{0,1,2,3,4,5,6,7,8,9})
	
	// Prints linked list nodes
	// current := ll1.head
	// for current != nil {
	// 	fmt.Println(current)
	// 	current = current.next
	// }

	// Add a cycle from tail to node 3
	targetNode := ll1.search(3)
	ll1.addCycle(targetNode)

	// Check to see if the cycle has been created
	// If it has, ll.tail.next should point to the node with value 3, and not nil
	// fmt.Println(ll1.tail.next.data)

	fmt.Println(containsCycle(&ll1))
	fmt.Println(containsCycle(&ll2))
}

// Takes pointer to linked list so you don't copy the whole linked list
func containsCycle(ll *LinkedList) bool {
	i := ll.head
	j := ll.head

	for j != nil {
		j = j.next.next
		i = i.next

		if i == j {
			return true
		}
	}
	return false
}
