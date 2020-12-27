package main

import "fmt"

const (
	currentyear = 2020
	a           = currentyear - iota
	b           = currentyear - iota
	c           = currentyear - iota
	d           = currentyear - iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
