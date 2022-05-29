package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("--------------")

	names := []string{"hiago", "joão", "matheus"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println("--------------")

	for key, value := range names {
		fmt.Println(key, ":", value)
	}

	fmt.Println("--------------")

	for _, value := range names {
		fmt.Println(value)
	}

	fmt.Println("--------------")

	var i int
	for i < len(names) {
		fmt.Println(names[i])
		i++
	}

	fmt.Println("--------------")

	for {
		fmt.Println("inside infinite loop")
		break
	}

	fmt.Println("--------------")

}
