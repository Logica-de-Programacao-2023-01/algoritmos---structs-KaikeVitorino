package main

import "fmt"

type Endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

func imprimirEnderecoCompleto(p Pessoa) {
	fmt.Printf("Endereço: %s, %d - %s, %s\n", p.endereco.rua, p.endereco.numero, p.endereco.cidade, p.endereco.estado)
}

func main() {
	endereco := Endereco{rua: "Rua Principal", numero: 123, cidade: "Cidade Exemplo", estado: "Estado Exemplo"}
	pessoa := Pessoa{nome: "Fulano", idade: 30, endereco: endereco}

	imprimirEnderecoCompleto(pessoa)
}
