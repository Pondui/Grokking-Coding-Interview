package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	graph := map[string][]string{
		"A": {"B","G"},
		"B": {"C", "D", "E"},
		"C": {},
		"D": {},
		"E": {"F"},
		"F": {},
		"G": {"H"},
		"H": {"I"},
		"I": {},
	}

	bfs(graph, "A")
	dfs(graph, "A")
}

// Use a deque to keep track of which nodes to visit
// Pop from the back, push from the front, i.e. FIFO
func bfs(graph map[string][]string, node string) {
	fmt.Println("BFS")
	var dq deque.Deque[string]
	visited := make(map[string]bool)

	// Initialise
	visited[node] = true
	dq.PushFront(node)

	for dq.Len() != 0 {
		current := dq.PopBack()
		fmt.Println(current)

		for _, e := range graph[current] {
			if _, ok := visited[e]; !ok {
				visited[e] = true
				dq.PushFront(e)
			}
		}

	}
}

// Use a deque to keep track of which nodes to visit
// Identical to bfs, except instead of visiting all nearest neighbours (FIFO)
// We visit one neighbour, then it's children
func dfs(graph map[string][]string, node string) {
	fmt.Println("DFS")
	var dq deque.Deque[string]
	visited := make(map[string]bool)

	// Initialise
	visited[node] = true
	dq.PushFront(node)

	for dq.Len() != 0 {
		current := dq.PopFront()
		fmt.Println(current)

		for _, e := range graph[current] {
			if _, ok := visited[e]; !ok {
				visited[e] = true
				dq.PushFront(e)
			}
		}

	}
}

