package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Hellooo James"},
	}

	for _, v := range x {
		for _, w := range v {
			fmt.Println(w)
		}
	}
	fmt.Println(x)
}
