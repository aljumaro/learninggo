package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong: ", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight: ", c.String())
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	var c1 Counter
	doUpdateWrong(c1)
	fmt.Println("in main: ", c1.String())
	doUpdateRight(&c1)
	fmt.Println("in main: ", c1.String())
}
