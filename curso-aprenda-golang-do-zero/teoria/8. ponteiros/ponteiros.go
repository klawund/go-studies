package main

import "fmt"

func main() {
	n1 := 10
	n2 := n1
	fmt.Println(n1, n2) // 10 10

	// altera só n1 porque o valor é passado por cópia, não referência
	n1++
	fmt.Println(n1, n2) // 11 10

	// para passar um valor por referência é necessário instanciá-lo como um ponteiro
	var n3 int
	var n4 *int // o * denota um ponteiro

	fmt.Println(n3, n4) // 0 <nil> - é nil pois o ponteiro não é primitivo

	n3 = 100
	// o ponteiro não pode receber um valor direto, apenas referências de memória
	// e a referência de memória é denotada pelo &
	n4 = &n3
	fmt.Println(n3, n4)  // 100 0xc0001a0e0 - o ponteiro recebe o endereço de memória, não o valor
	fmt.Println(n3, *n4) // 100 100 - isso faz a desreferenciação e traz o valor do endereço de memória apontado

	n3 = 150
	fmt.Println(n3, *n4) // 150 150 - depois de desreferenciado, o ponteiro traz o valor apontado em memória
}
