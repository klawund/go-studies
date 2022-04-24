package main

import (
	"fmt"
	"time"
)

func main() {
	go write("hello world") // go passa a execução do comando para outra thread
	write("coding go routines")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
