// Link to the problem: https://www.codechef.com/problems/CSUB

package main

import (
	"fmt"
	"strings"
)

func main() {
	var T uint32
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var N uint32
		fmt.Scan(&N)
		var str string
		fmt.Scan(&str)
		ones := strings.Count(str, "1")
		// Summation of total numbers of "1" in the string
		fmt.Println(ones * (ones + 1) / 2)
	}
}
