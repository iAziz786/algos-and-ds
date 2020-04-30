package singly

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func getStdoutLogs(fn func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return fmt.Sprintf("%s", out)
}

func linkNodes(length int) *Node {
	if length == 0 {
		return nil
	}
	node := Node{value: 1}
	for i, initial := 1, &node; i < length; i++ {
		nextNode := initial.Push(&Node{value: i + 1})
		initial = nextNode
	}
	return &node
}

func TestSingly(t *testing.T) {
	first := linkNodes(3)
	output := getStdoutLogs(first.Print)

	if output != "1\n2\n3\n" {
		t.Errorf("expected %s got %s", "1\n2\n3", output)
	}
}

func TestSinglyLinkedList_HeadIntersion(t *testing.T) {
	first := linkNodes(3)

	newHead := Node{value: 4, next: first}

	output := getStdoutLogs(newHead.Print)

	if output != "4\n1\n2\n3\n" {
		t.Fatal("failed to insert node at head")
	}
}

func TestSinglyLinkedList_Push(t *testing.T) {
	first := Node{value: 1, next: nil}
	second := first.Push(&Node{value: 2})
	_ = second.Push(&Node{value: 3})
	output := getStdoutLogs(first.Print)
	expectedOutput := "1\n2\n3\n"
	if output != expectedOutput {
		t.Errorf("Expected: %s\nGot: %s\n", output, expectedOutput)
	}
}
