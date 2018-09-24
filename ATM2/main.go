// Link to the problem: https://www.codechef.com/COOK98B/problems/ATM2

package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var N, K int
		fmt.Scanf("%d%d", &N, &K)
		A := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scanf("%d", &A[i])
		}
		Ans := ""
		for _, i := range A {
			if K-i >= 0 {
				Ans += "1"
				K -= i
			} else {
				Ans += "0"
			}
		}
		fmt.Println(Ans)
	}
}
