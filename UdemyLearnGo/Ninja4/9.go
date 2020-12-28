package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"}, // the []string here is not needed actually
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["Jack"] = []string{"programming", "drinking"}

	for k, v := range m {
		fmt.Printf("key=%v, ", k)
		for _, s := range v {
			fmt.Printf("%v, ", s)
		}
		fmt.Printf("\n")
	}
}
