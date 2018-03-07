// Link to the problem: https://www.codechef.com/problems/COMPILER

package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var INST string
		fmt.Scan(&INST)
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
