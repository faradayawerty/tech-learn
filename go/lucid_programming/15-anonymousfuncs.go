package main

import "fmt"

func main() {
	// anonymius function
	func(msg string) {
		fmt.Println(msg)
	}("Hello")
}
