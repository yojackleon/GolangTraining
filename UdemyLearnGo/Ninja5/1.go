package main

import "fmt"

type person struct {
	first        string
	last         string
	favIcecreams []string
}

func main() {
	p1 := person{
		first:        "Jack",
		last:         "Leon",
		favIcecreams: []string{"Peanut Butter", "chocolate", "pecan"},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.favIcecreams {
		fmt.Println(v)
	}

}
