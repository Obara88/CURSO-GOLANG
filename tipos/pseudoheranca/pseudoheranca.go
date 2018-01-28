package main

import (
	"fmt"
)

type carro struct {
	nome            string
	velocidadeAtual int
}

type banco struct {
	nome    string
	isCouro bool
}

type ferrari struct {
	carro
	nome        string
	turboLigado bool
	banco
}

func main() {

	f2 := ferrari{}

	f2.nome = "propriedade nome ferrari"
	f2.turboLigado = true
	f2.velocidadeAtual = 98 //propriedade do carro

	fmt.Println(f2)

	f2.carro.nome = "propriedade nome carro"
	f2.carro.velocidadeAtual = 99

	f2.banco.nome = "fabricante"
	f2.isCouro = true
	fmt.Println(f2)

}
