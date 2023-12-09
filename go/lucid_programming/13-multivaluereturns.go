package main

import "fmt"

func add_sub(a, b int) (int, int) {
	return a + b, a - b
}

func main() {

	ar, sr := add_sub(3, 2)
	fmt.Println("3 + 2 = ", ar)
	fmt.Println("3 - 2 = ", sr)
}
