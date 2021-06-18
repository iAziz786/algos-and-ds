package binary_search

// Compare check on two values if they both are equal it should return 0
// otherwise, if a is greather than b it should return 1 otherwise -1
type Compare = func(a, b interface{}) int

// Search look through the comparables and returns you the index which matches
// the provided value. If the item didn't match it will return -1.
func Search(array []interface{}, target interface{}, fn Compare) int {
	min, max := 0, len(array)-1
	for max > min {
		idx := midPoint(min, max)
		item := array[idx]

		if fn(item, target) < 0 {
			min = idx + 1
			continue
		}

		if fn(item, target) > 0 {
			max = idx - 1
			continue
		}

		return idx
	}
	return -1
}

func midPoint(min, max int) int {
	return min + (max-min)/2
}
