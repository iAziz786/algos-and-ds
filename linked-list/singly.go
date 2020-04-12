package singly

import "fmt"

type Node struct {
	value int
	next  *Node
}

func (n *Node) Print() {
	for ; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}
