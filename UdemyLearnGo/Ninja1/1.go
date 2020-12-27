package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", z)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
