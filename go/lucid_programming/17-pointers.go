package main

import "fmt"

func main() {
	val := 20
	fmt.Println(val)

	ptr := &val
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 10
	fmt.Println(*ptr)
	fmt.Println(val)

	val = 50
	fmt.Println(*ptr)
}
