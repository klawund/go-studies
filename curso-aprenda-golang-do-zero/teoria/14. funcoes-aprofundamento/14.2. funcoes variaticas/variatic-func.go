package main

import "fmt"

func sum(numbers ...int) int { // recebe vários parâmetros int e no final traz um slice para dentro da func
	var sum int

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func main() {
	sum := sum(1, 1, 1, 1, 1) // deve ser 5
	fmt.Println(sum)
}
