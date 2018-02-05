// Link to the problem https://www.codechef.com/problems/CNOTE

package main

import "fmt"

func main() {
	var T, X, Y, K, N, p, c int
	fmt.Scan(&T)
	for T > 0 {
		var flag int
		fmt.Scan(&X)
		fmt.Scan(&Y)
		fmt.Scan(&K)
		fmt.Scan(&N)
		for i := 0; i < N; i++ {
			fmt.Scan(&p)
			fmt.Scan(&c)
			if p >= (X-Y) && c <= K {
				flag = 1
			}
		}
		if flag > 0 {
			fmt.Println("LuckyChef")
		} else {
			fmt.Println("UnluckyChef")
		}
		T--
	}
}
