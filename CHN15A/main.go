// Link to the problem https://www.codechef.com/problems/CHN15A/

package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for T > 0 {
		var N, K, TRMGF int
		fmt.Scan(&N)
		fmt.Scan(&K)
		for N > 0 {
			var MINION int
			fmt.Scan(&MINION)
			if (MINION+K)%7 == 0 {
				TRMGF++
			}
			N--
		}
		fmt.Println(TRMGF)
		T--
	}
}
