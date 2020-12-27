package main

import "fmt"

func main() {
	i := 1975
	for {
		if i > 2020 {
			break
		}
		fmt.Println(i)
		i++
	}
}
