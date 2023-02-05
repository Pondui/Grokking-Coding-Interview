package main

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func newNode(data int, next, prev *Node) *Node {
	node := &Node{}
	node.data = data
	node.next = next
	node.prev = prev
	return node
}

func linkedListFromSlice(slc []int) LinkedList {
	var currNode *Node
	var prevNode *Node
	var ll LinkedList

	for i, e := range slc {
		// previous node in this run was the current node in the last run
		prevNode = currNode
		// instantiate new current node, and point it's prev field to the previous node
		currNode = newNode(e, nil, prevNode)
		// update the previous node's next field to the current node
		// can't do this on the 0th run since prevNode is not instantiated
		if i > 0 {
			prevNode.next = currNode
		}
		if i == 0 {
			ll.head = currNode
		} else if i == len(slc)-1 {
			ll.tail = currNode
		}
	}
	return ll
}
