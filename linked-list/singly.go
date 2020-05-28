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

// DeleteElementAtIndex will delete an element matching an index
func (node *Node) DeleteElementAtIndex(index int) bool {
	var prevNode *Node
	for i := 0; node != nil; i, prevNode, node = i+1, node, node.next {
		if node == nil {
			return false
		}
		// Get previous value
		if i == index {
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

// MakeMiddleNodeFirst finds the middle node and make it head of the list
// And return the middle node
func (node *Node) MakeMiddleNodeFirst() *Node {
	if node == nil || node.next == nil {
		return node
	}
	head := node
	middleNode, prevNode := head, head
	// find middle node
	for node != nil && node.next != nil {
		prevNode = middleNode
		middleNode = middleNode.next
		node = node.next.next
	}
	// make prevNode to point to next node of the middle node
	prevNode.next = prevNode.next.next
	// replace first node with middle node
	middleNode.next = head

	return middleNode
}

func moveNode(dest **Node, source **Node) {
	newNode := *source
	if newNode != nil {
		*source = newNode.next
		newNode.next = *dest
		*dest = newNode
	}
}

// MergeTwoSortedList will takes the head of two sorted list and
// returns the head newly merged array.
func MergeTwoSortedList(fSorted *Node, sSorted *Node) *Node {
	var result *Node = nil

	lastPtrRef := &result

	for {
		if fSorted == nil {
			*lastPtrRef = sSorted
			break
		} else if sSorted == nil {
			*lastPtrRef = fSorted
			break
		}
		if fSorted.value <= sSorted.value {
			moveNode(lastPtrRef, &fSorted)
		} else {
			moveNode(lastPtrRef, &sSorted)
		}
		lastPtrRef = &((*lastPtrRef).next)
	}
	return result
}
