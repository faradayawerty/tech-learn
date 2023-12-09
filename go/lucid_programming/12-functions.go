package main

import "fmt"

func main() {
	fmt.Println(add4(2, 3, 5, 7))
}

func add(a int, b int) int {
	return a + b
}

func add3(a, b, c int) int {
	return a + b + c
}

func add4(a, b, c, d int) int {
	return add(add(a, b), add(c, d))
}
