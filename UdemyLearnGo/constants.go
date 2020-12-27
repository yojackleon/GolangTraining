package main

import "fmt"

const x = 42
const y = "James Bond"
const z = true

const a int = 42
const b string = "James Bond"
const c bool = true

const (
	p = 42
	q = "James Bond"
	r = true
)
const (
	i int    = 42
	j string = "James Bond"
	k bool   = true
)

func main() {

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}
