package main

import "fmt"

const (
	_   = iota             // skip iota = 0
	ikb = 1 << (iota * 10) // 1 * 1000
	imb = 1 << (iota * 10) // 2 * 1000
	igb = 1 << (iota * 10) // 3 * 1000
)

func main() {
	x := 2
	fmt.Println("Decimal \t\t\t\t Binary")
	fmt.Printf("%d\t\t\t\t\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t\t\t\t\t%b\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t\t\t\t%b\n", ikb, ikb)
	fmt.Printf("%d\t\t\t\t\t%b\n", imb, imb)
	fmt.Printf("%d\t\t\t\t%b\n", igb, igb)

}
