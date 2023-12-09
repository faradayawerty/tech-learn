package main

import "fmt"

func main() {
	// go only has the for loop
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println()

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	fmt.Println()

	/* infinite loop */
	for {
		fmt.Println("Hello")
	}
}
