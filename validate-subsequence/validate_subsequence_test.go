package validate_subsequence

import (
	"testing"
)

func TestValidateSubsequence(t *testing.T) {
	var tests = []struct {
		message  string
		array    []int
		sequence []int
		match    bool
	}{
		{
			"should pass for single array",
			[]int{3},
			[]int{3},
			true,
		},
		{
			"should return false if the element is not present",
			[]int{3, 1},
			[]int{5},
			false,
		},
		{
			"should return true for both empty array and sequence",
			[]int{},
			[]int{},
			true,
		},
		{
			"should pass when order is correct",
			[]int{3, -45, 88, 12, -8},
			[]int{-45, 12, -8},
			true,
		},
		{
			"should break when order is not correct",
			[]int{-77, 45, -65, 100, 4},
			[]int{-65, 45, 100},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			didMatchCorrectly := ValidateSubsequence(tt.array, tt.sequence)
			if didMatchCorrectly != tt.match {
				t.Errorf("got %t, want %t", didMatchCorrectly, tt.match)
			}
		})
	}
}
