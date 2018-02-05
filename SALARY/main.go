// Link to problem: https://www.codechef.com/problems/SALARY

package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for T > 0 {
		var N, RESULT, wage, sumOfWages, SMALL int
		fmt.Scan(&N)
		for i := 0; i < N; i++ {
			fmt.Scan(&wage)
			sumOfWages += wage
			if i == 0 {
				SMALL = wage
			}

			if wage < SMALL {
				SMALL = wage
			}
		}
		/**
		* Instead of increasing the salary of all workers
		* Decrease from all the wages by minimun wage and
		* sum of the differences will be our RESULT
		*   W[1] - SMALL
		*	  W[2] - SMALL
		*   W[3] - SMALL
		*		...........
		*		...........
		*		SUM_OF_WAGES - N * SMALL
		 */
		RESULT = sumOfWages - N*SMALL
		fmt.Println(RESULT)
		T--
	}
}
