package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, value := range x {
		fmt.Printf("%v : type %T\n", value, value)
	}
}
