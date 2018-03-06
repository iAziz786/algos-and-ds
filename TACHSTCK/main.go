// Link to the problem: https://www.codechef.com/problems/TACHSTCK

package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, D, ans int
	fmt.Scan(&N)
	fmt.Scan(&D)
	L := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&L[i])
	}
	sort.Ints(L)
	if N <= 1 {
		fmt.Println(0)
	} else {
		for i := 1; i < N; {
			if L[i]-L[i-1] <= D {
				ans++
				i += 2
			} else {
				i++
			}
		}
		fmt.Println(ans)
	}
}
