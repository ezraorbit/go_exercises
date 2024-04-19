package main

import "fmt"

func foo(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	fmt.Println("I'm running after the program is done")
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
	defer foo([]int{1, 2, 3, 4, 5, 6}...)
	ages := bar([]int{1, 2, 3, 4, 5, 6})

	fmt.Println(ages)
	fmt.Println("I'm doing a okay!")
}
