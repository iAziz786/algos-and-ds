// Link to the problem: https://www.codechef.com/problems/COPS

package main

import (
	"fmt"
)

func main() {
	var T uint16
	fmt.Scan(&T)
	for T > 0 {
		var M, x, y int
		fmt.Scan(&M)
		fmt.Scan(&x)
		fmt.Scan(&y)
		var copInHouse, totalSafe int8
		isNotSafe := make([]bool, 100)
		for M > 0 {
			fmt.Scan(&copInHouse)
			var copMinRange, copMaxRange int16
			copMinRange = int16(copInHouse) - int16(x*y) - 1
			copMaxRange = int16(copInHouse) + int16(x*y) - 1
			if copMinRange < 0 {
				copMinRange = 0
			}
			if copMaxRange > 99 {
				copMaxRange = 99
			}
			for i := copMinRange; i <= copMaxRange; i++ {
				isNotSafe[i] = true
			}
			M--
		}
		for i := 0; i < 100; i++ {
			if isNotSafe[i] == false {
				totalSafe++
			}
		}
		fmt.Println(totalSafe)

		/**
		*	Below algorithm isn't made my me but this also works
		*	Comment in upper /\ code and comment out below \/ one
		* to see working
		 */

		// copInHouse := make([]int, M)
		// for i := 0; i < M; i++ {
		// 	fmt.Scan(&copInHouse[i])
		// }
		// sort.Ints(copInHouse)
		// if copInHouse[0]-x*y > 1 {
		// 	safeInitials := copInHouse[0] - x*y - 1
		// 	totalSafe = safeInitials
		// }
		// for i := 1; i < M; i++ {
		// 	safeInOther := copInHouse[i] - x*y - copInHouse[i-1] - x*y
		// 	if safeInOther > 0 {
		// 		k := safeInOther - 1
		// 		totalSafe += k
		// 	} else {
		// 		continue
		// 	}
		// }
		// if 100-copInHouse[M-1]-x*y > 0 {
		// 	h := 100 - copInHouse[M-1] - x*y
		// 	totalSafe += h
		// }
		// fmt.Println(totalSafe)

		T--
	}
}
