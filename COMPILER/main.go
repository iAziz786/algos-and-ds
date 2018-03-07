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
		fmt.Scanf("%s", &INST)
		var (
			count int
			res   int
		)
		for i, val := range INST {
			// ASCII for "<" = 60
			if val == 60 {
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
