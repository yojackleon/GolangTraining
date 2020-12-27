package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Printf("a is of type %T and has value %v\n", a, a)
	a = int(b)
	fmt.Println(a)
}
