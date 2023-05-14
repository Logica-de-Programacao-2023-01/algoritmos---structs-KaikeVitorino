package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func mediaNotas(a Aluno) float64 {
	soma := 0.0
	for _, nota := range a.notas {
		soma += nota
	}
	return soma / float64(len(a.notas))
}

func main() {
	a := Aluno{
		nome:  "João",
		idade: 20,
		notas: []float64{7.5, 8.0, 9.5},
	}
	media := mediaNotas(a)
	fmt.Printf("Média das notas do aluno %s: %.2f\n", a.nome, media)
}
