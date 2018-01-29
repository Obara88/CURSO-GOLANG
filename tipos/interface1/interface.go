package main

import (
	"fmt"
)

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(i imprimivel) {
	fmt.Println(i.toString())
}

func main() {
	var qqCoisa imprimivel = pessoa{"Rafael", "Obara"}

	imprimir(qqCoisa)

	qqCoisa = produto{"Batata", 1.99}

	imprimir(qqCoisa)

}
