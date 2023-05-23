package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func calcularArea(c Circulo) float64 {
	return math.Pi * c.raio * c.raio
}

func main() {
	c := Circulo{raio: 5}
	area := calcularArea(c)
	fmt.Printf("Área do círculo: %.2f\n", area)
}
