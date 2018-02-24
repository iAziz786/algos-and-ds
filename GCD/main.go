// No link to this problem just trying to implement Euclid’s GCD Algorithm

package main

import "fmt"

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var (
		num1 int
		num2 int
	)
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	if num1 != 0 || num2 != 0 {
		dividend := getMax(num1, num2)
		divider := getMin(num1, num2)
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
