// Link to the problem: https://www.codechef.com/MARCH18B/problems/CHEGLOVE

package main

import (
	"fmt"
)

func main() {
	var (
		T int8
	)
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var N int
		var (
			isFront = true
			isBack  = true
		)
		fmt.Scan(&N)
		L := make([]int, N)
		G := make([]int, N)
		// Scanning the fingers
		for i := 0; i < N; i++ {
			fmt.Scan(&L[i])
		}
		// Scanning the glove
		for i := 0; i < N; i++ {
			fmt.Scan(&G[i])
		}
		// Compare the length of the fingers
		for i, j := 0, N-1; i < N && j >= 0; i, j = i+1, j-1 {
			if L[i] > G[i] {
				isFront = false
			}
			if L[i] > G[j] {
				isBack = false
			}
		}
		// Checking the condition
		if isFront == true && isBack == true {
			fmt.Println("both")
		} else if isFront == true && isBack == false {
			fmt.Println("front")
		} else if isFront == false && isBack == true {
			fmt.Println("back")
		} else {
			fmt.Println("none")
		}
	}
}
