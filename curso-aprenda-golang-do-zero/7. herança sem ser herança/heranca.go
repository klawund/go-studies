package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// para fazer com que estudante "herde" os campos de pessoa, basta adicionar um campo
// que apenas seja o tipo do struct a ser herdado, sem um nome para o campo, assim todos
// os campos do struct pai são copiados para o struct filho
//
// isso é diferente de composição, para fazer composição bastaria fazer pessoa pessoa
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("herança")

	p1 := pessoa{"hiago", "klabunde", 19, 188}
	fmt.Println(p1)

	// não é possível fazer isso. é necessário instanciar o struct pai antes e passá-lo como argumento
	// e1 := estudante{"hiago", "klabunde", 19, 188, "ads", "senai"}

	// tecnicamente um struct estará dentro do outro, porém os campos serão acessados como
	// se estivessem apenas em um struct (e1.nome, e não e1.pessoa.nome)
	e1 := estudante{p1, "ads", "senai"}
	fmt.Println(e1)
}
