package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func main() {
	car1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Toyota",
		model: "Sedan",
		doors: 4,
		color: "Red",
	}
	car2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 4,
		color: "Green",
	}

	fmt.Println(car1.engine.electric, car1.model, car1.make, car1.doors, car1.color)
	fmt.Println(car2.engine.electric, car2.model, car2.make, car2.doors, car2.color)
	fmt.Printf("%#v\n", car1)
	fmt.Printf("%#v\n", car2)
}
