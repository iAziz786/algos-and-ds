package main

import "fmt"

// Short variable declarations are ONLY AVAILABLE WITHIN func
// So, := will not work here

func main() {
	var i, j int = 1, 2
	k := 3 // Short variable declaration with implicit type
	var c, python, java = true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
