package main

import "fmt"

func main() {
	x := 63848

	if x%2 == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")
	}

	y := 50
	if y < 50 {
		fmt.Println(y, "is less than 50")
	} else if y > 50 {
		fmt.Println(y, "is greater than 50")
	} else {
		fmt.Println(y, "is 50")
	}
}
