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
		last_name:                  "Ryanadale",
		favorite_ice_cream_flavors: []string{"Chocolate", "Vanilla"},
	}
	susan := person{
		first_name:                 "Susan",
		last_name:                  "Stormborn",
		favorite_ice_cream_flavors: []string{"Chocolate", "Vanilla", "Peanut Butter"},
	}

	last_name := map[string]person{
		duncan.last_name: duncan,
		susan.last_name:  susan,
	}

	for _, v := range last_name {
		fmt.Println(v.first_name)
		for _, flavors := range v.favorite_ice_cream_flavors {
			fmt.Printf("Favorite flavor: %v\n", flavors)
		}
		fmt.Println()
	}
}
