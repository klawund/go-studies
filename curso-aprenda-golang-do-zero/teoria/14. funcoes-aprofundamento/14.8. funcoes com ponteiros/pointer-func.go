package main

import "fmt"

func invertSignal(n int) int {
	return n * -1
}

func invertSignalWithPointer(n *int) {
	// como o que chega é um ponteiro, para utilizar o valor é necessário desreferenciar a variável
	*n = *n * -1
}

func main() {
	number := 20
	inverseNumber := invertSignal(number) // é passado uma cópia do valor de number - number permance inalterado
	fmt.Println(inverseNumber)

	number2 := 40
	invertSignalWithPointer(&number2) // como é recebido um ponteiro, é necessário passar como referência com &, não como valor
	fmt.Println(number2)
}
