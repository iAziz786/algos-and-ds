package main

import (
	"fmt"
)

func getMissingNumber(arr *[]int) int {
	a := 1
	b := (*arr)[0]

	for i := 2; i <= len(*arr)+1; i++ {
		a = a ^ i
	}

	for i := 1; i < len(*arr); i++ {
		b = b ^ (*arr)[i]
	}

	return b ^ a
}

func main() {
	a := []int{1, 2, 3, 4, 6, 7, 8}
	fmt.Println(getMissingNumber(&a))
}
