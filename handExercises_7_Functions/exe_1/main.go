package main

import "fmt"

func foo() int {
	return 73
}

func bar() (int, string) {
	return 19, "Stay happy"
}

func main() {
	age, saying := bar()
	fightingNumber := foo()

	fmt.Println(age, saying, fightingNumber)
}
