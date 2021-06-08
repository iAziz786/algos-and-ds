package queue_test

import (
	"testing"

	"github.com/iAziz786/algosandds/queue"
)

func TestEnqueue(t *testing.T) {
	q := queue.New()

	if q.Empty() != false {
		t.Errorf("queue should be empty")
	}

	q.Enqueue(1)
	q.Enqueue(11)
	q.Enqueue(1111)

	tests := []struct {
		want int
	}{
		{
			want: 1,
		},
		{
			want: 11,
		},
		{
			want: 1111,
		},
	}

	for _, tt := range tests {
		if res := q.Dequeue(); *res != tt.want {
			t.Errorf("dequeue expect %d, got %d", tt.want, res)
		}
	}

}
