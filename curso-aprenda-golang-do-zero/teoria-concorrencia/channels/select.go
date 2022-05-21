package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "channel 2"
		}
	}()

	for {
		messageChannel1 := <-c1
		fmt.Println(messageChannel1)

		messageChannel2 := <-c2
		fmt.Println(messageChannel2)
	}
}
