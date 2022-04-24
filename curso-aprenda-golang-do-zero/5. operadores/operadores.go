package main

import "fmt"

func main() {
	// aritméticos

	// os operadores aritméticos seguem o padrão de todas as langs
	// + - / * %

	num := 10 + 10 - 10*10/10%2
	fmt.Println(num)

	// atribuição

	// = para atribuição declarativa com tipo explícito e atribuição normal
	// := para atribuição declataiva com tipo implícito

	var nome string = "hiago"
	sobrenome := "klabunde"
	fmt.Println(nome, sobrenome)

	// relacionais

	// como a linguagem possui tipos estáticos não há operadores especiais para comparação estrita
	// > >= < <= == !=

	fmt.Println(1 > 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 < 1)
	fmt.Println(1 <= 1)
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)

	// lógicos

	// && || !

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// unários
	// ++ -- += -= *= /= %=

	num2 := 1
	num2++
	num2--
	num2 += 1
	num2 -= 1
	num2 *= 2
	num2 /= 2
	num2 %= 2
	fmt.Println(num2)

	// não há operador ternário!
}
