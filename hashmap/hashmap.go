package hashmap

import (
	"errors"

	"github.com/iAziz786/algosandds/linkedlist2"
)

type HashMap struct {
	internal []*linkedlist2.LinkedList
}

type KV struct {
	Key string
	Val string
}

var ErrNotFound = errors.New("not found")

const SIZE = 61 // using prime number for fewer possibility of collision

func New() *HashMap {
	return &HashMap{
		internal: make([]*linkedlist2.LinkedList, SIZE),
	}
}

func (h *HashMap) Insert(key string, val string) {
	idx := calculateIndex(key)
	kv := KV{Key: key, Val: val}

	// if it's the first entry at the index
	if h.internal[idx] == nil {
		h.internal[idx] = linkedlist2.New()
		h.internal[idx].PushBack(kv)
		return
	}

	h.internal[idx].PushBack(kv)
}

func (h *HashMap) Exists(key string) (found bool) {
	found = false

	// running the loop but instead there should have been a method to access
	// the value directly
	for i, m := 0, h.internal[calculateIndex(key)]; i < int(m.Size()); i++ {
		if val, err := m.ValueAt(uint64(i)); err == linkedlist2.ErrNotFound {
			continue
		} else {
			if kv, ok := val.(KV); ok && kv.Key == key {
				found = true
			}
		}
	}

	return
}

func (h *HashMap) Get(key string) (string, error) {
	var m = h.internal[calculateIndex(key)]

	// running the loop but instead there should have been a method to access
	// the value directly
	for i := 0; m != nil && i < int(m.Size()); i++ {
		if val, err := m.ValueAt(uint64(i)); err == linkedlist2.ErrNotFound {
			continue
		} else {
			if kv, ok := val.(KV); ok {
				return kv.Val, nil
			}
		}
	}

	return "", ErrNotFound
}

func (h *HashMap) Remove(key string) error {
	var m = h.internal[calculateIndex(key)]

	// running the loop but instead there should have been a method to access
	// the value directly
	for i := 0; m != nil && i < int(m.Size()); i++ {
		if _, err := m.ValueAt(uint64(i)); err == linkedlist2.ErrNotFound {
			continue
		} else {
			// remove the value at found index from the linked list
			m.Erase(uint64(i))
			return nil
		}
	}

	return ErrNotFound
}

// hash is a hash function which create hash of the string. It's an impl of
// Jenkins hash function https://en.wikipedia.org/wiki/Jenkins_hash_function
func hash(key string) uint32 {
	var hash uint32 = 0
	for _, char := range key {
		hash += uint32(char)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash
}

// calculateIndex find the position for the key
func calculateIndex(key string) int {
	return int(hash(key) % SIZE)
}
