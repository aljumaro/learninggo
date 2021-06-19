package main

import "fmt"

func main() {
	var x int = 10
	var y = 10
	var z int

	var a, b int = 10, 20
	var c, d int
	var e, f = 10, "hello"

	fmt.Println(x, y, z, a, b, c, d, e, f)

	var (
		g    int
		h        = 10
		i    int = 30
		j, k     = 40, "hello"
		l, m string
	)

	fmt.Println(g, h, i, j, k, l, m)

	n, o := 10, "hello"

	fmt.Println(n, o)
}
