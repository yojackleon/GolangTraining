package main

import "fmt"

func main() {
	name := "Jack"
	if name == "John" {
		fmt.Println("John")
	} else if name == "Jerry" {
		fmt.Println("Jerry")
	} else {
		fmt.Println("Jack")
	}
}
