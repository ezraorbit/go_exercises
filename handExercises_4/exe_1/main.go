package main

import "fmt"

func main() {
	/*
		UsingaCOMPOSITELITERAL:
			● create an ARRAY which holds 5 VALUES of TYPE int
			● assign VALUES to each index position.
		● Rangeover the array and print the values out.
			○ print out the VALUE and the TYPE
	*/
	x := [5]int{0, 1, 2, 3, 4}

	for _, value := range x {
		fmt.Printf("%v of type: %T\n", value, value)
	}
}
