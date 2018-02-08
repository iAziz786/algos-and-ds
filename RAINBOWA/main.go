// Link to the problem: https://www.codechef.com/problems/RAINBOWA

package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for T > 0 {
		var N, notEqual int
		rainbow := 1
		fmt.Scan(&N)
		A := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&A[i])
		}
		for i, j := 0, (N - 1); i < (N/2)+1; i, j = i+1, j-1 {
			if A[i] != A[j] {
				notEqual = 1
				break
			}
			if A[i] == rainbow {
				continue
			} else if A[i] == rainbow+1 {
				rainbow++
				continue
			} else {
				notEqual = 1
				break
			}
		}
		if notEqual == 1 || rainbow != 7 {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
		T--
	}
}
