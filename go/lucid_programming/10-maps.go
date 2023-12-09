package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 5
	m["b"] = 7

	fmt.Println(m)
	fmt.Println(m["a"])
	fmt.Println(len(m))

	delete(m, "a")
	fmt.Println(m)

	m2 := map[string]int{"one": 1, "two": 2}
	fmt.Println(m2)

	val, is_val_present := m["b"]
	fmt.Println(val)
	fmt.Println(is_val_present)
}
