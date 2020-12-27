package main

import "fmt"

func main() {
	//map[key]value{}
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println("For key Barnabas v=", v, "ok=", ok)

	if v, ok = m["Barnabas"]; ok {
		fmt.Println("yep exists")
	} else {
		fmt.Println("doesn't exist")
	}

	// adding a new key, value pair

	m["Jack"] = 45

	for k, v := range m {

		fmt.Println(k, v)
	}

	fmt.Println("Deleting")
	fmt.Println(m)
	delete(m, "Miss Moneypenny")

	fmt.Println(m)

	// deleting element that doesn't exist doesn't throw an error
	// so check it exists first.
	if v, ok = m["Ian Flemming"]; ok {
		delete(m, "Ian Flemming")
	}

}
