package main

import "fmt"

type employee struct {
	first_name string
	last_name  string
	id         int
}

func main() {
	e0 := employee{first_name: "Homer", last_name: "Simpson", id: 1}
	fmt.Println(e0)

	e1 := employee{"Michael", "Faraday", 228}
	fmt.Println(e1)

	// unmentioned values are zero-values
	e2 := employee{first_name: "Bob", id: 126}
	fmt.Println(e2)

	// accessing struct fields
	fmt.Println(e0.first_name)

	// accessing struct fields via pointer
	pe0 := &e0
	fmt.Println(pe0.first_name)

	// using pointers with structs
	fmt.Println(e0)
	fmt.Println(*pe0)
	pe0.first_name = "lol"
	fmt.Println(e0)
	fmt.Println(*pe0)

	// copying structs
	copy_e1 := e1
	fmt.Println(e1)
	fmt.Println(copy_e1)
	copy_e1.first_name = "lol"
	fmt.Println(e1)
	fmt.Println(copy_e1)
}
