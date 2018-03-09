// Link to the problem: https://www.codechef.com/problems/INTEST

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader

func readInt() int {
	var t int
	for i, _ := reader.ReadByte(); i >= '0' && i <= '9'; i, _ = reader.ReadByte() {
		t = t*10 + int(i-'0')
	}
	return t
}

func main() {
	var n, k, ans int
	reader = bufio.NewReader(os.Stdin)
	n, k = readInt(), readInt()
	for ; n > 0; n-- {
		var t int
		t = readInt()
		if t%k == 0 {
			ans++
		}
	}
	fmt.Println(ans)
}
