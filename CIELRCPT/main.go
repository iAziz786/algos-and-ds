// Link to the problem: https://www.codechef.com/problems/CIELRCPT

package main

import (
	"fmt"
)

func main() {
	var T uint8
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var (
			p   int
			ans int
		)
		fmt.Scan(&p)

		menu2 := [12]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

		var nextSmall int
		for p > 0 {
			for _, val := range menu2 {
				if p >= val {
					nextSmall = val
				} else {
					break
				}
			}
			ans++
			p -= nextSmall
		}

		fmt.Println(ans)
	}
}
