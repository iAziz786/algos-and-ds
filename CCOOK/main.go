// Link to the problem: https://www.codechef.com/problems/CCOOK

package main

import (
	"fmt"
)

func main() {
	var N uint16
	fmt.Scan(&N)
	for N > 0 {
		var sum uint8
		for i := 0; i < 5; i++ {
			var A uint8
			fmt.Scan(&A)
			sum += A
		}
		switch sum {
		case 0:
			fmt.Println("Beginner")
			break
		case 1:
			fmt.Println("Junior Developer")
			break
		case 2:
			fmt.Println("Middle Developer")
			break
		case 3:
			fmt.Println("Senior Developer")
			break
		case 4:
			fmt.Println("Hacker")
			break
		case 5:
			fmt.Println("Jeff Dean")
			break
		}
		N--
	}
}
