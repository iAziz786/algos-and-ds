// Link to the problem: https://www.codechef.com/problems/FRGTNLNG

package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for T > 0 {
		var N, K int
		fmt.Scan(&N)
		fmt.Scan(&K)
		dictOfFl := make([]string, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&dictOfFl[i])
		}
		isPresent := make([]bool, N)
		for K > 0 {
			var L int
			fmt.Scan(&L)
			var modernWord string
			for i := 0; i < L; i++ {
				fmt.Scan(&modernWord)
				for i, res := range dictOfFl {
					if isPresent[i] == false && res == modernWord {
						isPresent[i] = true
					}
				}
			}
			K--
		}
		for i := 0; i < N; i++ {
			if isPresent[i] == true {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
		T--
	}
}
