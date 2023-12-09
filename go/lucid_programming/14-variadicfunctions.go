package main

import "fmt"

func mult(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}

	return total
}

func main() {
	// Println is a variadic function
	fmt.Println("Hello", "world", "!!!")

	fmt.Println(mult(2, 3, 5, 7, 11, 13, 17))

	nums := []int{2, 3, 5, 7, 11}
	fmt.Println(mult(nums...))
}
