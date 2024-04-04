package main

import "fmt"

var name string = "Donald Duffy"

const hairColor string = "Black"

func main() {
	talkingSpeed := "really fast"
	fmt.Printf("%v has really cool %v hair\n", name, hairColor)
	fmt.Println(name + "talks " + talkingSpeed)
}
