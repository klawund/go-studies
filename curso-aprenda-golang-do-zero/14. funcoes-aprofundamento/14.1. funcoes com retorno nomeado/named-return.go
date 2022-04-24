package main

import "fmt"

func addAndSub(n1, n2 int) (sum, sub int) {
	// ambas as vars já estão declaradas, basta atribuir
	sum = n1 + n2
	sub = n1 - n2
	return // infere as duas variáveis nomeadas no retorno
}

func main() {
	sum, sub := addAndSub(10, 10)
	fmt.Println(sum, sub)
}
