package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for key, value := range m {
		fmt.Printf("key %v : \t%v\n", key, value)
	}
}
