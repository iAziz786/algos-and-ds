package singly

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSingly(t *testing.T) {
	first := Node{}
	second := Node{}
	third := Node{}

	third.value = 33
	third.next = nil

	second.value = 22
	second.next = &third

	first.value = 11
	first.next = &second

	// Start printing

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	first.Print()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	output := fmt.Sprintf("%s", out)
	if output != "11\n22\n33\n" {
		t.Errorf("expected %s got %s", "11\n22\n33", output)
	}
}
