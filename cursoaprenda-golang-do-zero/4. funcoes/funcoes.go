package main

import "fmt"

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// exemplo de arrow function
	var f = func(txt string) {
		fmt.Println("Função f")
		fmt.Println(txt)
	}

	f("texto da função")

	// para um retorno múltiplo, basta instanciar múltiplas variáveis com tipo inferido
	// a ordem importa, pois o retorno é transferido conforme a declaração da função
	soma2, sub2 := somaSubtracao(10, 20)
	fmt.Println(soma2, sub2)

	// caso não queira um dos retornos, use _
	soma3, _ := somaSubtracao(30, 40)
	fmt.Println(soma3)
}

// uma função pode ter múltiplos retornos
func somaSubtracao(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
