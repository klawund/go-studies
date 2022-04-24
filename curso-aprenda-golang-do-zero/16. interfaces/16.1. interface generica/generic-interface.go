package main

import "fmt"

// como a func recebe uma interface vazia como parâmetro, tecnicamente tudo "implementa" essa interface, assim, pode-se passar qualquer
// dado como parâmetro
func showData(i interface{}) {
	fmt.Println(i)
}

func main() {
	showData(123)
	showData("abc")
	showData(false)

	m := map[interface{}]interface{}{
		"name": 123,
		1.0:    false,
		true:   "klabunde",
	}
	fmt.Println(m)

	arr := []interface{}{"abc", 123, false}
	fmt.Println(arr)
}
