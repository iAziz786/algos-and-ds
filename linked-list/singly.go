package singly

import (
	"fmt"
)

// Node will be used to create new nodes in the list
type Node struct {
	value int
	next  *Node
}

// Print traverse the linked list and print each element
func (node *Node) Print() {
	for ; node != nil; node = node.next {
		fmt.Println(node.value)
	}
}

// Push will append the provided node at the end of the linked list
func (node *Node) Push(nodeToPush *Node) *Node {
	for ; ; node = node.next {
		// we have reached at the end node
		if node.next == nil {
			node.next = nodeToPush
			return nodeToPush
		}
	}
}

// DeleteKey the nodes with first matching key in the list
func (node *Node) DeleteKey(key int) bool {
	var prevNode *Node
	for ; node != nil; prevNode, node = node, node.next {
		if node == nil {
			return false
		}
		// Get previous value
		if node.value == key {
			// first element in the linked list
			if prevNode == nil {
				*node = *node.next
				return true
			}
			// last element
			if node.next == nil {
				prevNode.next = nil
				return true
			}
			prevNode.next = node.next
			return true
		}
	}
	return false
}
