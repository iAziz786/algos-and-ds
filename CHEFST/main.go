// Link to the problem: https://www.codechef.com/problems/CHEFST

package main

import (
	"bufio"
	"fmt"
	"os"
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
	var T, min int
	T = readInt()
	for ; T > 0; T-- {
		var n1, n2, m int
		n1 = readInt()
		n2 = readInt()
		if n1 < n2 {
			min = n1
		} else {
			min = n2
		}
		m = readInt()
		maxStoneOut := m * (m + 1) / 2
		if maxStoneOut < min {
			min = maxStoneOut
		}
		fmt.Println(n1 + n2 - (2 * min))
	}
}
