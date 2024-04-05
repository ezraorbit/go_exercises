package main

import "fmt"

/*
 create a for loop that uses break statement
 ● increment or decrement the variable being checked as a condition in the loop
*/

func main() {
	x := 0
	for {
		if x > 10 {
			break
		}

		fmt.Printf("x is currently %v\n", x)

		x++
	}
}
