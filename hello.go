package main

import "fmt"

// var for multiple variable declaration, last one is type
var c, python, java bool

func main() {
	// You can declare variable within function too
	var i int
	// Default for int = 0, and for bool = false
	fmt.Println(i, c, python, java)
}
