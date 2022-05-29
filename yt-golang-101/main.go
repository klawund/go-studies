package main

import (
	"fmt"
	"strconv"
)

func Hello(name string) {
	fmt.Println("hello,", name, "!")
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {
	i, _ := strconv.Atoi(b)
	if err != nil {
		return
	}

	total = a + i
	return
}

func main() {
	Hello("hiago")
	fmt.Println("total:", sum(10, 5))
	total, err := convertAndSum(10, "5")
	fmt.Println("total:", total, err)
}
