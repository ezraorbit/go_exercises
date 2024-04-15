package main

import "fmt"

func main() {
	steven := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Steven",
		friends: map[string]int{
			"Jilian":    4,
			"Sammantha": 2,
		},
		favDrinks: []string{"Coffee", "Milk", "Cocacola"},
	}

	fmt.Println(steven)
	fmt.Printf("%#v\n", steven)
}
