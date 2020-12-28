package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type family struct {
	mum      person
	dad      person
	children []person
	language string
}

func main() {
	p1 := person{
		first: "Jack",
		last:  "Leon",
		age:   45,
	}
	p2 := person{
		first: "Carolina",
		last:  "Leon",
		age:   43,
	}
	p3 := person{"Jake", "Leon", 16}
	p4 := person{"Josh", "Leon", 14}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fam1 := family{
		dad:      p1,
		mum:      p2,
		children: []person{p3, p4},
		language: "Spanish",
	}
	fmt.Println(fam1)

}
