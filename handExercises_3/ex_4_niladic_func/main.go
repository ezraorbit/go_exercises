package main

import (
	"fmt"
	"math/rand"
)

/*
 Modify the previous program to have your program use the init func to print
 ○ "This is where initialization for my program occurs
*/

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {

	//create a random int between 0 - 250
	var x int = rand.Intn(251)

	// print out the name and value of the variables
	fmt.Printf("x has a value of %v\n", x)

	switch {
	case x < 101:
		fmt.Println("Between 0 and 100")
	case x < 201:
		fmt.Println("Between 101 and 200")
	default:
		fmt.Println("Between 201 and 250")
	}

}
