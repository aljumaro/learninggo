package main

import "fmt"

func main() {
	//Slices are not comparable

	var x = []int{10, 20, 30}

	fmt.Println(x)

	var y []int //Zero value: nil

	fmt.Println(y == nil) //true

	y = append(y, 10)

	fmt.Println(y)

	z := []int{20, 30, 40}
	y = append(y, z...)

	fmt.Println(y)
}
