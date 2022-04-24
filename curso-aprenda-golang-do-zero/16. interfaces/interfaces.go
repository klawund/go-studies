package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func showArea(s shape) {
	fmt.Printf("A área da forma é %0.2f\n", s.area())
}

func main() {
	r := rectangle{10, 15}
	showArea(r)

	c := circle{5.2}
	showArea(c)
}
