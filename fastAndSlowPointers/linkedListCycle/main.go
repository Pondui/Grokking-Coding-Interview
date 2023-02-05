package main

import "fmt"

func main() {
	ll := linkedListFromSlice([]int{0,1,2,3,4,5})
	fmt.Println("Head:", ll.head)
	fmt.Println("Tail:", ll.tail)
}
