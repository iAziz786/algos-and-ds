// Link to the problem: https://www.codechef.com/problems/MAXDIFF

package main

import (
	"fmt"
	"sort"
)

func main() {
	var T uint8
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var K, N uint8
		var dadSum, sonSum int
		fmt.Scan(&N)
		fmt.Scan(&K)
		W := make([]int, N)
		for i := uint8(0); i < N; i++ {
			fmt.Scan(&W[i])
		}
		sort.Ints(W)
		if K > N/2 {
			for i := N - 1; i >= N-K; i-- {
				dadSum += W[i]
			}
			for i := uint8(0); i < N-K; i++ {
				sonSum += W[i]
			}
		} else {
			for i := uint8(0); i < K; i++ {
				sonSum += W[i]
			}
			for i := K; i < N; i++ {
				dadSum += W[i]
			}
		}
		fmt.Println(dadSum - sonSum)
	}
}
