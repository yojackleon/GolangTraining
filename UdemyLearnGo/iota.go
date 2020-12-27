package main

import "fmt"

const (
	a = iota + 1
	b
	c
)
const (
	x = iota
	y
	z
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
