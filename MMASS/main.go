// Link to the problem: http://www.spoj.com/problems/MMASS/

package main

import (
	"fmt"
)

func main() {
	var (
		chem string
		x    int
		res  int
	)
	fmt.Scan(&chem)
	stack := []int{}
	for _, val := range chem {
		strVal := string(val)
		if strVal == "(" {
			stack = append(stack, -1)
		} else if strVal == "H" {
			stack = append(stack, 1)
		} else if strVal == "C" {
			stack = append(stack, 12)
		} else if strVal == "O" {
			stack = append(stack, 16)
		} else if strVal >= "2" && strVal <= "9" {
			x, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, (x * int(val-48)))
		} else if strVal == ")" {
			j := len(stack) - 1
			var total int
			for stack[j] != -1 {
				x, stack = stack[len(stack)-1], stack[:len(stack)-1]
				total += x
				j--
			}
			x, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, total)
		}
	}
	for _, val := range stack {
		res += val
	}
	fmt.Println(res)
}
