package main

import "fmt"

func main() {
	fmt.Println("arrays e slices")
	// arrays

	// assim como em java, o arr tem tamanho fixo
	var arr1 [5]int
	arr1[0] = 10
	fmt.Println(arr1)

	arr2 := [2]string{"hiago", "klabunde"}
	fmt.Println(arr2)

	// ... infere o tamanho do arr a partir da qtd de elementos
	arr3 := [...]int{1, 2, 3, 4}
	fmt.Println(arr3)

	// slices

	// o slice é um arr com tamanho flexível
	slice := []int{10, 11}
	slice = append(slice, 18) // o mesmo que o .push() do js - retorna um novo slice com o valor passado no final
	fmt.Println(slice)

	slice2 := arr1[1:3] // o 1 será inclusivo e o 3 exclusivo
	fmt.Println(slice2)

	arr1[1] = 15 // altera o arr1 e o slice2 pois há um ponteiro para o arr1

	// arrays internos

	// o slice sempre vai ser uma fatia de um array existente em memória;
	// quando o slice não é criado a partir de um array existente, ele irá criar um array interno próprio;
	// o array interno do slice tem uma capacidade, porém quando essa capacidade é estourada
	// ela é dobrada automaticamente, o que faz com que o slice não tenha um tamanho limite

	slice3 := make( // função que aloca memória para um objeto novo
		[]float32, // tipo do objeto que será instanciado
		10,        // tamanho do slice retornado
		11)        // capacidade máxima do array interno do slice

	fmt.Println(len(slice3)) // tamanho do slice - 10
	fmt.Println(cap(slice3)) // capacidade do slice - 11

	slice3 = append(slice3, 0.1)
	fmt.Println(cap(slice3)) // 22 - como estourou com o append, dobrou a capacidade
}
