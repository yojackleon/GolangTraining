package main

import "fmt"

func main() {
	x := 7
	switch x {
	case 2, 7, 8: // multiple cases, but you can't repeat cases conditions.
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
		fallthrough // will print 5 eventhough x is not 5 - WTF ?
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")

	default:
		fmt.Println("default")
	}

	switch {

	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}
