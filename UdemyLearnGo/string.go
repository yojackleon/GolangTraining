package main

import (
	"fmt"
)

func main() {

	s := "Hello Jack"
	fmt.Println(s)

	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	byteSlice := []byte(s)

	fmt.Println(byteSlice)
	fmt.Println(byteSlice[0])

	for i := 0; i < len(byteSlice); i++ {
		fmt.Printf("%#U", byteSlice[i])
	}
	fmt.Println()
	for i, v := range s {
		fmt.Println(i, v)
	}

}
