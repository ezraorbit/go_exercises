package main

import "fmt"

func main() {
	x := 98

	fmt.Printf("Memory address: %v\nType: %T\nData stored at memory address: %v\n", &x, &x, x)
}
