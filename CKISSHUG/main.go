// Problem: https://www.codechef.com/problems/CKISSHUG
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MOD is max iteration
var MOD int64 = 1000000007

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	inputs := strings.Fields(text)
	if len(inputs) == 0 {
		return
	}
	T, _ := strconv.ParseInt(inputs[0], 10, 64)
	for ; T > 0; T-- {
		text, _ := reader.ReadString('\n')
		inputs = strings.Fields(text)
		N, _ := strconv.ParseInt(inputs[0], 10, 64)
		fmt.Println(calcAns(N))
	}
}

func pow(a, b int64) int64 {
	ans := int64(1)
	for b > 1 {
		if b%2 != 0 {
			ans *= a
			ans %= MOD
		}
		a *= a
		b /= 2
		if a > MOD {
			a %= MOD
		}
	}
	return ((ans * a) % MOD)
}

func sumGP(n int64) int64 {
	return 2 * (pow(2, n) - 1) % MOD
}

func calcAns(n int64) int64 {
	if n%2 == 0 {
		return (sumGP(n/2) + pow(2, n/2)) % MOD
	}
	return sumGP((n + 1) / 2)
}
