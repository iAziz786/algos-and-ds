// Link to the problem: https://www.codechef.com/MARCH18B/problems/MIXCOLOR

package main

import (
	"fmt"
)

func main() {
	var T int8
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var N, ans int
		m := make(map[int]int)
		fmt.Scan(&N)
		for i := 0; i < N; i++ {
			var A int
			fmt.Scan(&A)
			m[A]++
		}
		for _, v := range m {
			ans += (v - 1)
		}
		fmt.Println(ans)
	}
}
