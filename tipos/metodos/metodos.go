package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomecompleto string) {
	partes := strings.Split(nomecompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	//na pratica deveria passar o endereço de p1: (&p1).setNomeCompleto("Maria Pereira")
	//mas o Go quebra esse galho pra vc...
	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto())
}
