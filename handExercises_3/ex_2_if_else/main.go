package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//create a random int between 0 - 250
	var x int = rand.Intn(251)

	// print out the name and value of the variables
	fmt.Printf("x has a value of %v\n", x)

	if x < 101 {
		fmt.Println("Between 0 and 100")
	} else if x < 201 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}

}
