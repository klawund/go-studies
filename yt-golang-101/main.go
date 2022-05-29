package main

import "fmt"

var (
	name string
	n1   int
	n2   int32
)

func main() {
	msg := "aula 101 - 03"
	fmt.Println(msg)

	var b0, f, s = true, 2.3, "olá"
	fmt.Println(b0, f, s)

	var a, b, c int32
	fmt.Println(a, b, c)

	var total int
	total++
	fmt.Println("total:", total)

	name = "hiago"
	fmt.Println("hello,", name, "!")

	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
