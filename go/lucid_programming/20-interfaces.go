package main

import "fmt"
import m "math"

// define an interface
type geometry interface {
	area() float64
	perim() float64
}

// define rect and circle
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// implement area for rect and circle
func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return m.Pi * c.radius * c.radius
}

// implement perim for rect and circle
func (r rect) perim() float64 {
	return 2*(r.width + r.height)
}

func (c circle) perim() float64 {
	return 2 * m.Pi * c.radius
}

// example of the interface application
func mesure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	mesure(r)
	mesure(c)
}
