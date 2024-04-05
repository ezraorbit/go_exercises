package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Create 2 random ints between 0 inclusive and 10 exclusive
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("This is the value of x:\t %v\n", x)
	fmt.Printf("This is the value of y:\t %v\n", y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y < 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("none of the previous cases were met")
	}
}
