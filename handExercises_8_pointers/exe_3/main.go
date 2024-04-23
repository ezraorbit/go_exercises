package main

import "fmt"

type human struct {
	first string
}

func main() {
	p := human{
		first: "Groq",
	}

	fmt.Println(p)

	p = changeName(p, "Jen")
	fmt.Println(p)

	changeNamePtr(&p, "JMo")
	fmt.Println(p)
}

func changeName(h human, s string) human {
	h.first = s
	return h
}

func changeNamePtr(h *human, s string) {
	h.first = s
}
