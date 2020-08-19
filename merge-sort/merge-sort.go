package merge_sort

import (
	"math"
)

// Merge left and right arrays
// start - start index of the array
// middle - middle index of the array
// end - end index of the array
func Merge(input []int, start, middle, end int) {
	itemsAtLeftSide := middle - start + 1
	itemsAtRightSide := end - middle

	// creating arrays which holds one extra element then the
	// currently needed because of the sentinel assignment
	left := make([]int, itemsAtLeftSide+1)
	right := make([]int, itemsAtRightSide+1)

	// allocate elements in the left side of the array
	for i := 0; i < itemsAtLeftSide; i++ {
		left[i] = input[start+i]
	}

	// allocate elements in the right side of the array
	for j := 0; j < itemsAtRightSide; j++ {
		// adding one because because j index is staring from zero
		// which in most of the cases conflict with the value in
		// the left array, adding one removes the conflict element
		right[j] = input[middle+j+1]
	}

	// sentinel assignments
	left[itemsAtLeftSide] = math.MaxInt64
	right[itemsAtRightSide] = math.MaxInt64

	for i, j, k := 0, 0, start; k <= end; k++ {
		if left[i] <= right[j] {
			// assign what is at the top of the left array
			input[k] = left[i]
			i++
		} else {
			// assign what is at the top of the right array
			input[k] = right[j]
			j++
		}
	}
}

func MergeSort(input []int, start, end int) {
	if start < end {
		// getting the middle value
		middle := (start + end) / 2

		// sort whatever in the left of the array
		MergeSort(input, start, middle)

		// sort whatever in the right of the array
		MergeSort(input, middle+1, end)

		// merge left and right sorted values
		Merge(input, start, middle, end)
	}
}
