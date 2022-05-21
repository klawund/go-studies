package main

import (
	"fmt"
	"time"
)

func main() {
	// o canal é uma espécie de deque, que recebe infos dos dois lados (main routine e goroutine)
	// e serve como um sincronizador de rotinas
	channel := make(chan string)

	go write("hello world", channel)

	msg := <-channel
	fmt.Println(msg)
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}
}
