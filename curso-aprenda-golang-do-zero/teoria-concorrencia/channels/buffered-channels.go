package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Olá mundo"

	message := <-channel
	fmt.Println(message)
}
