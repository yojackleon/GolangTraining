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

	p2 := person{
		first:        "Carolina",
		last:         "Leon",
		favIcecreams: []string{"Coffee", "chocolate", "strawberry"},
	}

	m := map[string]person{
		"dad": p1,
		"mum": p2,
	}

	for k, v := range m {
		fmt.Println(k)
		for _, fav := range v.favIcecreams {
			fmt.Println(fav)
		}
	}

}
