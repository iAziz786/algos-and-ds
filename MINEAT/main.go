// Link to the problem: https://www.codechef.com/problems/MINEAT

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var reader *bufio.Reader

func readInt() int {
	var sign, n int = 1, 0
	i, _ := reader.ReadByte()
	for ; i < '0' || i > '9'; i, _ = reader.ReadByte() {
		if i == '-' {
			sign = -1
		}
	}
	for ; i >= '0' && i <= '9'; i, _ = reader.ReadByte() {
		n = n*10 + int(i-'0')
	}
	return (n * sign)
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	T := readInt()
	for ; T > 0; T-- {
		N, H := readInt(), readInt()
		A := make([]int, N)
		for i := 0; i < N; i++ {
			A[i] = readInt()
		}
		sort.Ints(A)

		left, right := 1, A[N-1]
		for left < right {
			mid := int(math.Floor(float64(left+right) / float64(2)))
			hoursTaken := 0
			for _, v := range A {
				hoursTaken += int(math.Ceil(float64(v) / float64(mid)))
			}
			if hoursTaken <= H {
				right = mid
			} else if hoursTaken > H {
				left = mid + 1
			}
		}
		fmt.Println(left)
	}
}
