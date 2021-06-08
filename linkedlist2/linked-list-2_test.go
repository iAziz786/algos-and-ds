package linkedlist2_test

import (
	"testing"

	"github.com/iAziz786/algosandds/linkedlist2"
)

func TestLinkedList_NewLinkedList(t *testing.T) {
	ll := linkedlist2.New()

	if !ll.Empty() {
		t.Errorf("size != %t but %t", true, ll.Empty())
	}
}

func TestLinkedList_ValueAt(t *testing.T) {
	ll := linkedlist2.New()

	ll.PushFront(42)
	ll.PushFront(43)

	if val, _ := ll.ValueAt(1); val != 42 {
		t.Errorf("want %d, got %d", 42, val)
	}

	if val, _ := ll.ValueAt(0); val != 43 {
		t.Errorf("want %d, got %d", 43, val)
	}

	if _, err := ll.ValueAt(2); /* means out of bound */ err != linkedlist2.ErrNotFound {
		t.Errorf("out of bound should return error\n")
	}
}

func TestLinkedList_PushFront(t *testing.T) {
	ll := linkedlist2.New()

	ll.PushFront(42)

	if ll.Size() != 1 {
		t.Errorf("size != %d but %d\n", 1, ll.Size())
	}
}

func TestLinkedList_PopFront(t *testing.T) {
	ll := linkedlist2.New()

	if err := ll.PopFront(); err == nil {
		t.Errorf("empty list pop should return error\n")
	}

	ll.PushFront(99)
	ll.PushFront(101)

	if err := ll.PopFront(); err != nil {
		t.Errorf("can't pop error because %s", err)
	}

	if ll.Size() != 1 {
		t.Errorf("error size: want=%d, got=%d", 1, ll.Size())
	}

	if val, _ := ll.ValueAt(0); val != 99 {
		t.Errorf("popping should removing front")
	}
}

func TestLinkedList_PushBack(t *testing.T) {
	ll := linkedlist2.New()

	ll.PushBack(87)

	if ll.Size() != 1 {
		t.Errorf("want size = %d, got %d", 1, ll.Size())
	}

	if val, _ := ll.ValueAt(0); val != 87 {
		t.Errorf("value must be %d", 87)
	}

	ll.PushBack(88)

	if ll.Size() != 2 {
		t.Errorf("want size = %d, got %d", 1, ll.Size())
	}

	if val, _ := ll.ValueAt(1); val != 88 {
		t.Errorf("value must be %d", 88)
	}
}

func TestLinkedList_PopBack(t *testing.T) {
	ll := linkedlist2.New()

	ll.PopBack()

	if ll.Size() != 0 {
		t.Errorf("popping on empty list should have no op")
	}

	ll.PushFront(97)
	ll.PushFront(101)
	ll.PushFront(103)

	ll.PopBack()

	if ll.Size() != 2 {
		t.Errorf("error size: want=%d, got=%d", 2, ll.Size())
	}

	if _, err := ll.ValueAt(2); err == nil {
		t.Errorf("popping should remove from back")
	}

	ll.PopBack()

	if ll.Size() != 1 {
		t.Errorf("error size: want=%d, got=%d", 1, ll.Size())
	}

	if _, err := ll.ValueAt(1); err == nil {
		t.Errorf("popping should remove from back")
	}
}

func TestLinkedList_InsertFirst(t *testing.T) {
	ll := linkedlist2.New()

	ll.Insert(0, 1)

	if val, err := ll.ValueAt(0); err != nil || val != 1 {
		t.Errorf("inserting at index %d should include %d instead return %d", 0, 1, val)
	}
}

func TestLinkedList_InsertMiddle(t *testing.T) {
	ll := linkedlist2.New()

	ll.PushFront(3)
	ll.PushFront(1)
	ll.PushFront(0)

	ll.Insert(2, 2)

	if val, err := ll.ValueAt(2); err != nil || val != 2 {
		t.Errorf("inserting at index %d should include %d", 2, 2)
	}
	if val, err := ll.ValueAt(3); err != nil || val != 3 {
		t.Errorf("inserting at index should move later values")
	}
}

func TestLinkedList_InsertLast(t *testing.T) {
	ll := linkedlist2.New()

	ll.PushFront(1)
	ll.Insert(1, 2)

	if val, err := ll.ValueAt(1); err != nil || val != 2 {
		t.Errorf("inserting at index %d should include %d", 1, 2)
	}
}

func TestLinkedList_ReverseEmpty(t *testing.T) {
	ll := linkedlist2.New()

	ll.Reverse()

	if ll.Size() != 0 {
		t.Errorf("reversing should not empty the list")
	}
}

func TestLinkedList_ReverseSingle(t *testing.T) {
	ll := linkedlist2.New()

	ll.PushFront(1)
	ll.Reverse()

	if val, _ := ll.ValueAt(0); val != 1 {
		t.Errorf("single element reverse should be no ops")
	}
}

func TestLinkedList_ReverseOddElements(t *testing.T) {
	ll := linkedlist2.New()

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)
	ll.Reverse()

	if val, _ := ll.ValueAt(1); val != 2 {
		t.Errorf("middle element in the odd list should not change")
	}

	if val, _ := ll.ValueAt(0); val != 3 {
		t.Errorf("should switch last element with first")
	}

	if val, _ := ll.ValueAt(2); val != 1 {
		t.Errorf("should switch first element with last")
	}
}

func TestLinkedList_ReverseEvenElements(t *testing.T) {
	ll := linkedlist2.New()

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)
	ll.PushBack(4)
	ll.Reverse()

	tests := []struct {
		idx uint64
		val int
	}{
		{
			idx: 0,
			val: 4,
		},
		{
			idx: 1,
			val: 3,
		},
		{
			idx: 2,
			val: 2,
		},
		{
			idx: 3,
			val: 1,
		},
	}

	for _, tt := range tests {
		if val, _ := ll.ValueAt(tt.idx); val != tt.val {
			t.Errorf("after reversting found index %d should contain %d", tt.idx, tt.val)
		}
	}
}
