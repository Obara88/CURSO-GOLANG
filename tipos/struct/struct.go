package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Metodo : não mais é do que uma função com receiver
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.12,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	var produto2 produto
	produto2 = produto{"Lapis", 1.12, 0.1}
	fmt.Println(produto2, produto1.precoComDesconto())

	produto3 := produto{"Lapis", 1.12, 0.99}
	fmt.Println(produto3, produto1.precoComDesconto())

}
