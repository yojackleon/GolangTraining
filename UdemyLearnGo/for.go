package main

import "fmt"

func main() {
	var i, j int
	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {

			fmt.Println(i, ",", j)

		}
	}
	i = 0
	for i < j {
		fmt.Println(i)
		i++
	}
	for {
		i--
		if i < 5 {
			break
		}
		if i == 7 {
			continue
		}
		fmt.Println(i)

	}
}
