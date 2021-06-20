package main

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	//do something here
	fmt.Println(opts)
	return nil
}

func addTo(base int, values ...int) []int {
	out := make([]int, 0, len(values))
	for _, v := range values {
		out = append(out, base+v)
	}
	return out
}

func main() {
	err := MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      15,
	})

	fmt.Println(err)

	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
}
