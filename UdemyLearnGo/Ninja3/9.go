package main

import "fmt"

func main() {
	favSport := "Running"

	switch favSport {
	case "Running":
		fmt.Println("Hell Yeah")
	case "Anything else":
		fmt.Println("Not so much ")
	default:
		fmt.Println("Default")
	}

}
