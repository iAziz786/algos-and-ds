// Link ot the problem: https://www.codechef.com/SEPT18B/problems/CHEFADV

package main

import (
	"fmt"
)

// CHEFADV is solving the problem at https://www.codechef.com/SEPT18B/problems/CHEFADV
func CHEFADV() string {
	var N, M, X, Y int
	knowladge, power := 1, 1
	fmt.Scanf("%d%d%d%d", &N, &M, &X, &Y)
	// N = Knowladge, M = Power
	N -= knowladge
	M -= power
	if (N%X == 0) && (M%Y == 0) {
		return "Chefirnemo"
	} else if (N >= 1 && M >= 1) && ((N-1)%X == 0) && ((M-1)%Y == 0) {
		return "Chefirnemo"
	}
	return "Pofik"
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		res := CHEFADV()
		fmt.Println(res)
	}
}
