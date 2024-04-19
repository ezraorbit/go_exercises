package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hey there my name is %s, and I'm %d years old.\n", p.first, p.age)
}

func main() {
	p1 := person{
		first: "Sandra",
		age:   23,
	}

	p1.speak()
}
