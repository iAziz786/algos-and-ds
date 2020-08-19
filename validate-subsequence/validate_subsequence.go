package validate_subsequence

// ValidateSubsequence will receive two arrays, one the main and
// other a subsequence which wanted to match whether it's present
// in the first array or not
func ValidateSubsequence(array []int, subsequence []int) bool {
	matchCount := 0
	for idxSubseq, idxArr := 0, 0; idxSubseq < len(subsequence) && idxArr < len(array); {
		if subsequence[idxSubseq] == array[idxArr] {
			matchCount++
			idxSubseq++
			continue
		}
		idxArr++
	}

	if matchCount == len(subsequence) {
		return true
	}
	return false
}
