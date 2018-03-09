package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func readStr() string {
	str := []string{}
	i, _ := reader.ReadByte()
	for ; i != '\n'; i, _ = reader.ReadByte() {
		str = append(str, string(i))
	}
	return strings.Join(str, "")
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	T := readStr()
	fmt.Printf("%T %v\n", T, T)
}
