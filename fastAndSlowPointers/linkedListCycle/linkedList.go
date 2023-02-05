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

// O(1)
func (l *LinkedList) insertNodeAtStart(val int) {
	newNode := &Node{
		data: val,
		next: l.head,
		prev: nil,
	}
	l.head.prev = newNode
	l.head = newNode
}

// O(1)
func (l *LinkedList) insertNodeByPtr(val int, node *Node) {
	newNode := &Node{
		data: val,
		next: node,
		prev: node.prev,
	}
	node.prev.next = newNode
	node.prev = newNode
}

// O(1)
func (l *LinkedList) deleteNodeByPtr(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

// O(1), adds cycle starting from tail so some node
// Also possible to start from any node, but all nodes after are no longer part of the linked list
func (l *LinkedList) addCycle(node *Node) {
	l.tail.next = node
}

// O(n) since all nodes must be visited in the worst case scenario
func (l *LinkedList) search(target int) *Node {
	current := l.head
	// last node's next pointer is nil
	for current != nil {
		if current.data == target {
			return current
		}
		current = current.next
	}
	return nil
}
