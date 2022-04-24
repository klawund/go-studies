package main

import "fmt"

// a função init sempre é chamada antes da main;
// cada arquivo pode possuir a sua função init, que é executada sempre que o arquivo vai ser utilizado
func init() {
	fmt.Println("função init sendo executada")
}

func main() {
	fmt.Println("função main sendo executada")
}
