package main

import "fmt"

func main() {
	func(text string) {
		fmt.Println(text)
	}("hello world") // IIFE
}
