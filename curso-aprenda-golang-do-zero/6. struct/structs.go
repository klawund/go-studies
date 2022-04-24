package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

type usuarioComEndereco struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("structs.go")

	// declaração explícita
	var u usuario
	u.nome = "hiago"
	u.idade = 19
	fmt.Println(u)

	// declaração implícita com todos os campos do struct
	u2 := usuario{"hiago", 19}
	fmt.Println(u2)

	// declaração implícita com named arguments
	u3 := usuario{nome: "hiago"}
	u4 := usuario{idade: 19}
	fmt.Println(u3, u4)

	// declaração de struct compost (struct com struct dentro)
	endereco := endereco{"rua 123", 123}
	u5 := usuarioComEndereco{"hiago", 19, endereco}
	fmt.Println(u5)
}
