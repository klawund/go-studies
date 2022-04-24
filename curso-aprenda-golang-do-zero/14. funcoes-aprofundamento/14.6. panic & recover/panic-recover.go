package main

import "fmt"

// o panic() serve para interromper o funcionamento do programa (mata o processo)
// é usado em conjunto com o defer e com o recover()
// o recover() tenta recuperar a execução do programa após um panic e é chamado via defer

func recoverExecution() {
	// convenção:
	// if init com variável "r" que recebe o resultado de recover()
	// se r != null é porque foi possível recuperar a exexcução
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func calculateAvg(n1, n2 float32) bool {
	defer recoverExecution() // é executado antes dos returns e panic; nos returns não fará nada pois recover() vai retornar nil
	avg := (n1 + n2) / 2

	// nesse caso o programa entrará em pânico caso a média seja igual a 6
	if avg < 6 {
		return false
	} else if avg > 6 {
		return true
	}

	panic("A MÉDIA É IGUAL A 6")
}

func main() {
	avg := calculateAvg(6, 6)
	fmt.Println(avg)
	fmt.Println("Execução finalizada")
}
