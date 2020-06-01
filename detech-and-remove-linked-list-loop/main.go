package main

import "fmt"

// Node is a node in a linked list
type Node struct {
	value int
	next  *Node
}

func hashAndRemove(head *Node) {
	if head == nil || head.next == nil {
		return
	}

	slow, fast := head, head

	slow = slow.next
	fast = fast.next.next

	for fast != nil && fast.next != nil {
		if slow == fast {
			break
		}
		slow = slow.next
		fast = fast.next.next
	}

	if slow == fast {
		slow = head
		for slow.next != fast.next {
			slow = slow.next
			fast = fast.next
		}
		fast.next = nil
	}
}

func printList(head *Node) {
	for ; head != nil; head = head.next {
		fmt.Println(head.value)
	}
}

func main() {
	head := &Node{50, nil}
	head.next = &Node{20, nil}
	head.next.next = &Node{15, nil}
	head.next.next.next = &Node{4, nil}
	head.next.next.next.next = &Node{10, nil}

	head.next.next.next.next.next = head.next.next

	hashAndRemove(head)

	printList(head)
}
