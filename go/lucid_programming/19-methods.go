package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{10, 5}
	fmt.Println("area:", r.area())

	r_ptr := &r
	fmt.Println("perim:", r_ptr.perim())

}
