// Problem: https://www.geeksforgeeks.org/find-subarray-with-given-sum/
package main

import "fmt"

func main() {
	arr := []int{1, 4, 0, 0, 3, 10, 5}
	subarr := []int{}
	currentSum := 0
	expectedSum := 7
	for _, val := range arr {
		for currentSum > expectedSum {
			currentSum -= subarr[0]
			subarr = subarr[1:]
		}
		if currentSum == expectedSum {
			break
		}
		subarr = append(subarr, val)
		currentSum += val
	}
	fmt.Println(subarr)
}
