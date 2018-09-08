// Link to the problem: https://www.codechef.com/SEPT18B/problems/MAGICHF
package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var N, X, S int
		fmt.Scanf("%d%d%d", &N, &X, &S)
		for ; S > 0; S-- {
			var s1, s2 int
			fmt.Scanf("%d%d", &s1, &s2)
			if s1 == X {
				X = s2
			} else if s2 == X {
				X = s1
			}
		}
		fmt.Println(X)
	}
}
