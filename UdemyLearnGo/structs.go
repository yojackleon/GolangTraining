package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	p   person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   23,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   43,
	}
	p3 := person{"Jack", "Leon", 45}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	sa1 := secretAgent{
		p3,
		true,
	}
	fmt.Println(sa1)

	sa2 := secretAgent{
		p:   person{"Todd", "McLeod", 50},
		ltk: false,
	}
	fmt.Println(sa2.p)
}
