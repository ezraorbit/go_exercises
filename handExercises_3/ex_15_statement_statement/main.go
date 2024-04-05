package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if x := rand.Intn(5); x == 3 {
		for i := 0; i < 100; i++ {
			fmt.Println("x is 3")
		}
	} else {
		fmt.Println("Maybe next time champ🥲")
	}
}
