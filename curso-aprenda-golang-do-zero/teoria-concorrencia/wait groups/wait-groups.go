package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // thread pool

	waitGroup.Add(2) // número de threads

	go func() {
		write("hello world")
		waitGroup.Done() // diz para o pool que a thread finalizou
	}()

	go func() {
		write("coding go routines")
		waitGroup.Done()
	}()

	// até ter recebido o número de dones equivalente ao número de threads, segura a exec do programa
	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
