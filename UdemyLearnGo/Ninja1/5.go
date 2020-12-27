package main

import "fmt"

var a int

type newint int

var x newint
var y int

func main() {

	x = 43
	y = 44
	fmt.Printf("x is of type %T and has value %v\n", x, x)
	y = int(x)

	fmt.Printf("y is of type %T and has value %v\n", y, y)

}
