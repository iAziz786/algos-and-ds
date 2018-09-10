// Link to the problem: https://www.codechef.com/LTIME63B/problems/EID

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader

func readInt() int {
	var t int
	for i, _ := reader.ReadByte(); i >= '0' && i <= '9'; i, _ = reader.ReadByte() {
		t = t*10 + int(i-'0')
	}
	return t
}

// EID return the result
func EID(a []int) int {
	sort.Ints(a)
	size := len(a)
	diff := a[size-1]
	for i := 0; i <= size-2; i++ {
		tempDiff := a[i+1] - a[i]
		if tempDiff < diff {
			diff = tempDiff
		}
	}
	return diff
}

func main() {
	var T int
	reader = bufio.NewReader(os.Stdin)
	T = readInt()
	for ; T > 0; T-- {
		var N int
		N = readInt()
		A := make([]int, N)
		for i := 0; i < N; i++ {
			A[i] = readInt()
		}
		res := EID(A)
		fmt.Println(res)
	}
}
