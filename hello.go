package main

import (
	"fmt"
)

// When declaring a variable without explicity specifying it's type
// Its type is inferred by the value on the right hand side

func main() {
	v := "Hello"

	fmt.Printf("v is the typeof %T\n", v)
}
