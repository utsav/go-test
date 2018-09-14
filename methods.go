package main

import "fmt"

type rect struct {
	width int
	height int
}

func (c *rect) area() int {
	return c.width * c.height;
}

func main() {
	r := rect{width: 7, height: 4}

	fmt.Println(r.area());
}