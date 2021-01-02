package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Print("(", sum, ",")
		sum += x
		fmt.Print(sum, ")\n")
		return sum
	}

}
func main() {

	pos, _ := adder(), adder()

	for i := 0; i < 5; i++ {

		fmt.Print("i=", i)
		fmt.Print("pos(i)=", pos(i))
		fmt.Print("\n-----\n")
	}

}
