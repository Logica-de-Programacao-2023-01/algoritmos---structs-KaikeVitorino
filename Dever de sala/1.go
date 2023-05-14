package main

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func calculaArea(r Retangulo) float64 {
	return r.largura * r.altura
}

func main() {
	r := Retangulo{largura: 5, altura: 7}
	area := calculaArea(r)
	fmt.Println("Área do retângulo:", area)
}
