package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)

	//*
	//[] 0 0
	//[1] 1 1
	//[1 1] 2 2
	//[1 1 1] 3 4
	//[1 1 1 1] 4 4
	//[1 1 1 1 1] 5 8
	//*/

	var y = make([]int, 0, 10)
	fmt.Println(y, len(y), cap(y))
}
