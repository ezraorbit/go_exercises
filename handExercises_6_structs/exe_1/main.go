package main

import "fmt"

type person struct {
	first_name                 string
	last_name                  string
	favorite_ice_cream_flavors []string
}

func main() {
	duncan := person{
		first_name:                 "Duncan",
		last_name:                  "Stormborn",
		favorite_ice_cream_flavors: []string{"Chocolate", "Vanilla"},
	}
	susan := person{
		first_name:                 "Susan",
		last_name:                  "Stormborn",
		favorite_ice_cream_flavors: []string{"Chocolate", "Vanilla", "Peanut Butter"},
	}

	fmt.Printf("%#v", duncan)
	fmt.Printf("%#v", susan)

	for key, value := range susan.favorite_ice_cream_flavors {
		fmt.Println(key, value)
	}
}
