package hashmap_test

import (
	"testing"

	"github.com/iAziz786/algosandds/hashmap"
)

func TestHashMap_Insert_Get(t *testing.T) {
	h := hashmap.New()

	h.Insert("billie", "lovely")

	val, err := h.Get("billie")

	if err != nil {
		t.Error(err)
		return
	}

	if val != "lovely" {
		t.Errorf("want %s, got %s", "lovely", val)
	}
}

func TestHashMap_Not_Present(t *testing.T) {
	h := hashmap.New()

	_, err := h.Get("billie")

	if err != hashmap.ErrNotFound {
		t.Errorf("if key is not present it should return error")
	}
}

func TestHashMap_Contains(t *testing.T) {
	h := hashmap.New()

	h.Insert("billie", "your power")

	if !h.Exists("billie") {
		t.Errorf("key %s should be present", "billie")
	}
}

func TestHashMap_Remove(t *testing.T) {
	h := hashmap.New()

	h.Insert("billie", "your power")
	h.Insert("ed", "i don't care")
	h.Insert("khalid", "young dump and broke")

	if err := h.Remove("ed"); err != nil {
		t.Error(err)
	}

	if _, err := h.Get("ed"); err != hashmap.ErrNotFound {
		t.Errorf("deleting should remove the key and value")
	}
}
