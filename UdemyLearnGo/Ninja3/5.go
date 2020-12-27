package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("when i=%v then %v mod 4=%v \n", i, i, i%4)
	}
}
