// No link to this problem just trying to implement Euclid’s GCD Algorithm

package main

import "fmt"

func getMaxMin(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func main() {
	var (
		num1 int
		num2 int
	)
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	if num1 != 0 || num2 != 0 {
		dividend, divider := getMaxMin(num1, num2)
		reminder := dividend % divider
		for reminder > 0 {
			temp := divider
			divider = reminder
			dividend = temp
			reminder = dividend % divider
		}
		fmt.Println(divider)
	} else {
		fmt.Println(0)
	}
}
