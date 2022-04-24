package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8
	// int16
	// int32
	// int64
	var numero0 int8 = 12
	fmt.Println(numero0)

	// inferido pela arquitetura do pc (32 ou 64)
	var numero1 int = -100
	fmt.Println(numero1)

	// unassigned int
	var numero2 uint32 = 2
	fmt.Println(numero2)

	// rune == int32
	var numero3 rune = 12345
	fmt.Println(numero3)

	// byte == int8
	var numero4 byte = 1
	fmt.Println(numero4)

	// float32
	// float64
	var real1 float32 = 1234.56
	fmt.Println(real1)

	// infere pela arquitetura do pc, igual o int
	real3 := 123456.78
	fmt.Println(real3)

	var str string = "essa é uma string"
	fmt.Println(str)

	str2 := "essa é outra string"
	fmt.Println(str2)

	// não existe char, apenas rune/byte que recebe o valor a partir de um string literal com aspas simples
	char := 'B'       // infere como rune/int8
	fmt.Println(char) // 66 (equivalente à posição na tabela ascii)

	// todos os tipos primitivos possuem valor padrão
	// string = ""
	// intx = 0
	// floatx = 0.0
	// error = nil
	// bool = false

	var bool1 bool
	fmt.Println(bool1)

	// para instanciar um novo erro é necessário usar o pacote interno "errors"
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
