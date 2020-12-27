package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 0
	x[1] = 1
	x[2] = 2
	x[3] = 3
	x[4] = 4
	var y = [5]int{6, 7, 8, 9, 10}
	fmt.Println(x)
	fmt.Println(y)

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T", x)
}
