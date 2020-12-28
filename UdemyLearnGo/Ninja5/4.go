package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

func main() {
	v := struct {
		doors int
		color string
	}{
		doors: 2,
		color: "black",
	}
	fmt.Println(v)
}
