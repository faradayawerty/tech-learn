package main

import "fmt"

func main() {
	i := 4

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("something else")
	}

	j := 13

	switch {
	case j == 7:
		fmt.Println(j, "is equal to 7")
	case j < 7:
		fmt.Println(j, "is less than 7")
	case j > 7:
		fmt.Println(j, "is greater than 7")
	}
}
