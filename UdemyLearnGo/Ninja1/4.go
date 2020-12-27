package main

import "fmt"

var a int

type newint int

var x newint

func main() {

	x = 43
	fmt.Printf("x is of type %T and has value %v\n", x, x)

}
