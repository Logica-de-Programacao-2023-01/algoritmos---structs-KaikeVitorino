package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func calcularArea(t Triangulo) float64 {
	area := t.base * t.altura / 2
	return area
}

func main() {
	triangulo := Triangulo{base: 5, altura: 3}
	area := calcularArea(triangulo)
	fmt.Println("Área do triângulo:", area)
}
