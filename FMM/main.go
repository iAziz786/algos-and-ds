// This is implementation of Fast Module Multiplication

package main

import (
	"fmt"
)

func fastExp(base, exp int) uint {
	const MOD = 10000000000000007
	res := 1
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}
	return uint(res % MOD)
}

func main() {
	var (
		x int
		y int
	)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Println(fastExp(x, y))
}
