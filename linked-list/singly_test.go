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

func TestSingly(t *testing.T) {
	first := Node{value: 11}
	second := Node{value: 22}
	third := Node{value: 33}

	first.next = &second
	second.next = &third
	// Start printing

	output := getStdoutLogs(first.Print)

	if output != "11\n22\n33\n" {
		t.Errorf("expected %s got %s", "11\n22\n33", output)
	}
}

func TestSinglyLinkedList_HeadIntersion(t *testing.T) {
	first := Node{value: 11}
	second := Node{value: 22}
	third := Node{value: 33}

	third.value = 33
	third.next = nil

	second.value = 22
	second.next = &third

	first.value = 11
	first.next = &second

	newHead := Node{value: 55, next: &first}

	output := getStdoutLogs(newHead.Print)

	if output != "55\n11\n22\n33\n" {
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
