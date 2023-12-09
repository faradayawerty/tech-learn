package main

import "fmt"

func main() {
	const pi float64 = 3.14159265
	fmt.Println(pi)

	const n int = 9e9
	fmt.Println(n)

	fmt.Println(2+7i)
	fmt.Println(010)
	fmt.Println(0x100)
	fmt.Println(0o10)
	fmt.Println(0b11010)
}
