package main

import "fmt"

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "I'm 008."}
	data := [][]string{jb, mp}
	for i, v1 := range data {
		fmt.Println(v1[i])
		for j, word := range v1 {
			fmt.Println(j, word)
		}
	}
	fmt.Println(data)
}
