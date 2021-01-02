package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int { // same type, no need to repeat type
	return a + b + c
}

func vals() (int, int) { //multiple returns
	return 3, 7

}

func sum(nums ...int) int { //variadic function
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)

	fmt.Println("6+7 = ", plus(6, 7))

	a, b := vals()
	_, c := vals() // don't need the first returned value

	fmt.Println(a, b, c)

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	as := []int{6, 7, 8, 9, 45, 23, 12}

	fmt.Println(sum(as...)) // how to use a slice to call a variadic function

}
