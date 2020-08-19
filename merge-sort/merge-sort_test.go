package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		message string
		in      []int
		out     []int
	}{
		{
			"should return single value when input is single",
			[]int{3},
			[]int{3},
		},
		{
			"should sort values in random order",
			[]int{3, 1},
			[]int{1, 3},
		},
		{
			"should return empty array if the input is also empty",
			[]int{},
			[]int{},
		},
		{
			"should sort values with duplication",
			[]int{3, 2, 3},
			[]int{2, 3, 3},
		},
		{
			"should sort values with negative cases",
			[]int{-9, 2, 3, -3},
			[]int{-9, -3, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			MergeSort(tt.in, 0, len(tt.in)-1)
			if reflect.DeepEqual(tt.in, tt.out) != true {
				t.Errorf("got %q, want %q", tt.in, tt.out)
			}
		})
	}
}
