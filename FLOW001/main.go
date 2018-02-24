// Link to the problem: https://www.codechef.com/problems/FLOW001

package main

import (
	"fmt"
)

func main() {
	var T uint16
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var num1, num2 int
		fmt.Scan(&num1)
		fmt.Scan(&num2)
		fmt.Println(num1 + num2)
	}
}
