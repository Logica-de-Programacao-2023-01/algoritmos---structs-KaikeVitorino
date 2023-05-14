package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
	Ano    int
}

func infos(L Livro) {
	fmt.Printf("Titulo: %s\n", L.titulo)
	fmt.Printf("Autor: %s\n", L.autor)
	fmt.Printf("Ano de publicação: %d\n", L.Ano)
}

func main() {
	Livro := Livro{"Harry Potter e a Pedra Filosofal", "J. K. Rowling", 1997}
	infos(Livro)
}
