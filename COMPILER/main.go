// Link to the problem: https://www.codechef.com/problems/COMPILER

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

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
	var T int
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var INST string
		fmt.Fscan(reader, &INST)
		var (
			count int
			res   int
		)
		for i, val := range INST {
			if string(val) == "<" {
				count++
			} else {
				count--
			}
			if count == 0 {
				res = i + 1
			} else if count < 0 {
				break
			}
		}
		fmt.Println(res)
	}
}
