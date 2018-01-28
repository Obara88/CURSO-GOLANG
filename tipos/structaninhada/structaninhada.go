package main

import (
	"fmt"
)

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {

	var resultado float64

	for _, item := range p.itens {
		resultado += item.preco * float64(item.qtde)
	}

	return resultado
}

func main() {
	pedido1 := pedido{
		userID: 0,
		itens: []item{
			item{produtoID: 0, qtde: 1, preco: 10},
			item{produtoID: 0, qtde: 1, preco: 20},
			item{produtoID: 0, qtde: 1, preco: 30},
		},
	}

	fmt.Println(pedido1.valorTotal())
}
