package main

import "fmt"

// You can include initializer in var declaration
var i, j int = 1, 2

func main() {
	// If initializer in present, variable gets it's type upon value
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
