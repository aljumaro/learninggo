package main

import "fmt"

func main() {

	var x [3]int
	var y = [3]int{10, 20, 30}
	fmt.Println(x, y)

	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	//[1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
	fmt.Println(z)

	var a = [...]int{10, 20, 30}
	fmt.Println(a)

	fmt.Println(a == y) //true

	fmt.Println(len(z))
}
