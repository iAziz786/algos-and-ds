// Link to the problem: https://www.codechef.com/MARCH18B/problems/BIGSALE

package main

import (
	"fmt"
)

func main() {
	var T int8
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var N int
		var totalLoss float64
		fmt.Scan(&N)
		for i := 0; i < N; i++ {
			pqd := make([]int, 3)
			var incPrice, disPrice float64
			for j := 0; j < 3; j++ {
				fmt.Scan(&pqd[j])
			}
			incPrice = (float64(pqd[0]) + float64(pqd[0]*pqd[2])/100)
			disPrice = (incPrice - (incPrice*float64(pqd[2]))/100)
			totalLoss += (float64(pqd[1]) * (float64(pqd[0]) - disPrice))
		}
		fmt.Println(totalLoss)
	}
}
