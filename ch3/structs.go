package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	fred := person{
		"fred",
		38,
		"roger",
	}
	fmt.Println(fred)

	julia := person{
		name: "julia",
		age:  45,
	}

	julia.pet = "toby"
	fmt.Println(julia.name, julia.pet)
}
