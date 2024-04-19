package main

import "fmt"

func foo(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}

	return sum
}

func bar(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func main() {
	ages := bar([]int{1, 2, 3, 4, 5, 6})
	fighters := foo([]int{1, 2, 3, 4, 5, 6}...)

	fmt.Println(ages, fighters)
}
