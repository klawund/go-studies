package main

import "fmt"

// o método é uma função que sempre estará associada a um tipo - struct ou interface

type user struct {
	name  string
	idade uint8
}

// passando o tipo antes do nome, a função será atrelada a esse tipo, e pode usar os seus campos
// o tipo é inferido via parâmetro, então há uma variável "u" de tipo user
// u fica disponível dentro do método e user fica inferido como o tipo atrelado ao método
// se o tipo for passado sem ponteiro, o objeto do struct é passado via cópia e o objeto original fica inalterado
func (u user) save() {
	fmt.Printf("salvando os dados de %s no banco\n", u.name)
}

// como user é passado via ponteiro, esse método consegue alterar os dados reais do objeto que o acessou
func (u *user) alterName(newName string) {
	u.name = newName
}

func main() {
	u := user{"hiago", 19}
	u.save()

	u.alterName("klabunde")
	fmt.Println(u.name)
}
