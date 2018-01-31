package main

import (
	"fmt"

	"github.com/Obara88/goarea/subpacote"

	"github.com/Obara88/goarea"
)

func main() {
	//para instalar o pacote, use o comando
	//go get -u github.com/Obara88/goarea
	fmt.Println(goarea.Circ(4.0))
	subpacote.Hello()
}
