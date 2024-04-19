package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func main() {
	s1 := square{
		length: 10,
		width:  10,
	}

	c1 := circle{
		radius: 10,
	}

	info(s1)
	info(c1)
}

func info(s shape) {
	fmt.Printf("The area of the shape is %.2f\n", s.area())
}
