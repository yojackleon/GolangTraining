package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	v         vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		v: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(t1.v.color)
	fmt.Println(s1.color)
	fmt.Println(s1)
}
