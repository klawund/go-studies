package main

import "fmt"

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func dayOfWeek2(number int) string {
	// é possível fazer o switch dessa forma, e criar uma condição para cada caso
	switch {
	case number == 1:
		return "Domingo"
	default:
		return "Número inválido"
	}
}

func dayOfWeek3(number int) string {
	var day string

	switch number {
	case 1:
		day = "Domingo"
		fallthrough // faz com que o case não possua um break, e siga para o próximo caso (sem reavaliar a condição) após a sua execução
	case 2:
		day = "Segunda-feira"
	default:
		day = "Número inválido"
	}

	return day
}

func main() {
	day := dayOfWeek(3)
	fmt.Println(day)

	day = dayOfWeek2(1)
	fmt.Println(day)

	day = dayOfWeek3(1) // como tem o fallthrough irá retornar Segunda-feira ao invés de domingo
	fmt.Println(day)
}
