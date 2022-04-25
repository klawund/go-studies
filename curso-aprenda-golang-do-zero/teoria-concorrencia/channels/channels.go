package main

import (
	"fmt"
	"time"
)

// o channel é um tipo de comunicação entre a thread main e a thread da goroutine;
// o channel pode receber e enviar informações;

func main() {
	// todo channel tem que ter um tipo de dado definido
	channel := make(chan string)
	go write("hello world", channel)

	fmt.Println("executed after calling write()")

	// passa a informação do canal para a variável
	// irá segurar a execução até que o canal receba a informação vinda da goroutine
	message := <-channel
	fmt.Println(message)
}

func write(text string, channel chan string) {
	// passa o valor da variável para o canal
	channel <- text
	time.Sleep(time.Second * 2)
}
