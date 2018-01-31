// Link to the problem: https://www.codechef.com/problems/LECANDY

package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N, C int
		fmt.Scan(&N)
		fmt.Scan(&C)
		A := make([]int, N)
		// Scanning the Array candies
		for j := 0; j < N; j++ {
			fmt.Scan(&A[j])
		}
		var sum int
		for _, s := range A {
			sum += s
		}
		if C >= sum {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
