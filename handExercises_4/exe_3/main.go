package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{}
	z := []int{}
	a := []int{}
	b := []int{}
	b = append(b, x[:5]...)
	y = append(y, x[5:]...)
	z = append(z, x[2:7]...)
	a = append(a, x[1:6]...)

	fmt.Println(b)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
}
