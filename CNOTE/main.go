// Link to the problem https://www.codechef.com/problems/CNOTE

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
	var T int
	T = readInt()
	for ; T > 0; T-- {
		var X, Y, K, N int
		X = readInt()
		Y = readInt()
		K = readInt()
		N = readInt()
		var lucky = false

		for ; N > 0; N-- {
			var p, c int
			p = readInt()
			c = readInt()
			if lucky != true && p >= (X-Y) && c <= K {
				lucky = true
			}
		}

		if lucky == true {
			fmt.Println("LuckyChef")
		} else {
			fmt.Println("UnluckyChef")
		}
	}
}
