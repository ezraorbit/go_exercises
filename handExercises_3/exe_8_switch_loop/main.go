package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("The loop is beginning➰")
	for i := 0; i < 42; i++ {
		fmt.Printf("Counter is on: %v\n", i)
		newNum()
	}

}

func newNum() {
	x := rand.Intn(5)

	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	case 4:
		fmt.Println("x is 4")
	default:
		fmt.Println("Well this is embarassing🥲")
	}
}
