package main

import "fmt"

func main() {
	//composite literal
	x := []int{4, 5, 7, 8, 42}
	// a slice allows you to group together values of the SAME type
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x[1:3])
	fmt.Println(x[2:])
	fmt.Println(x[:2])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	y := []int{56, 3, 4, 2, 4}
	x = append(x, y...) // appending slices
	fmt.Println(x)

	x = append(x, 56, 23, 12) // adding numbers to slices
	fmt.Println(x)

	// x = append(x, 12, y...) doesn't work !!

	x = append(x[:2], x[8:]...) // deleting elements from slice
	fmt.Println(x)

	// arrays are the underlying type of slices
	// when a slice grows a new array is created - that can be wasteful
	// so we can great a slice with enough capacity for it to grow
	// using make( []type, initialsize, capacity)
	z := make([]int, 10, 100)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	// but you can't assign values outside of the range ( ie len)
	// you have to append
	// if you append past the capacity, the runtime just doubles the capacity

}
