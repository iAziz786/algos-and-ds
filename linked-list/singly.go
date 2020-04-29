package singly

import "fmt"

// Node will be used to create new nodes in the list
type Node struct {
	value int
	next  *Node
}

// Print traverse the linked list and print each element
func (n *Node) Print() {
	for ; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}
