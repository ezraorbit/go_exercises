package main

import (
	"fmt"
	"log"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for key, value := range m {
		fmt.Printf("key %v : \t%v\n", key, value)
	}

	age := m["James"]
	fmt.Println(age)

	if age, ok := m["Q"]; ok {
		fmt.Println(age)
	} else {
		log.Println("Yoh that doesn't exist🤦‍♂️")
	}
}
