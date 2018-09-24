// Link to the problem: https://www.codechef.com/COOK98B/problems/MAKPERM

package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var N int
		fmt.Scanf("%d", &N)
		A := make([]int, N)
		mask := make([]bool, N)
		for i := 0; i < N; i++ {
			fmt.Scanf("%d", &A[i])
			mask[i] = false
		}
		for _, val := range A {
			if val <= N && !mask[val-1] {
				mask[val-1] = true
			}
		}
		ans := 0
		for _, val := range mask {
			if !val {
				ans++
			}
		}
		fmt.Println(ans)
	}
}
