package main

import (
	"fmt"
)

// sem defer
func isStudentApproved(grades ...float32) bool {
	fmt.Println("Entrando na função de verificação de aprovação do aluno.")
	var sum float32

	for _, grade := range grades {
		sum += grade
	}

	avg := sum / float32(len(grades))

	if avg >= 6 {
		fmt.Println("Média calculada. Resultado será retonado")
		return true
	}

	fmt.Println("Média calculado. Resultado será retornado")
	return false
}

// com defer
func isStudentApproved2(grades ...float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função de verificação de aprovação do aluno.")
	var sum float32

	for _, grade := range grades {
		sum += grade
	}

	avg := sum / float32(len(grades))

	if avg >= 6 {
		// defer será executado aqui
		return true
	}

	// defer será executado aqui
	return false
}

func main() {
	r1 := isStudentApproved(10, 10, 10)
	r2 := isStudentApproved2(5, 5, 5)
	fmt.Println(r1, r2)
}
