package main

import "fmt"

func main() {
	str_slice := []string{"a", "b", "c", "d"}

	for i, c := range str_slice {
		fmt.Println("index", i)
		fmt.Println("character", c)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println("key:", k, "val:", v)
	}

	for _, w := range m {
		fmt.Println("value:", w)
	}
}
