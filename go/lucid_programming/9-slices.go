package main

import "fmt"

func main() {
	var s []int  // slice (variable length)
	var a [5]int // array (fixed length)

	s = make([]int, 3)

	// similar to arrays
	a[0] = 4
	a[1] = 8
	a[2] = 16
	fmt.Println(a)

	s[0] = 4
	s[1] = 8
	s[2] = 16
	fmt.Println(s)

	// append function is unique to slices
	s = append(s, 32, 64)
	fmt.Println(s)

	// slice syntax
	fmt.Println(s[1:3])

	u := make([][]int, 3)
	fmt.Println(u)
}
