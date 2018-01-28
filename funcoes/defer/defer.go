package main

import (
	"fmt"
)

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim1")
	defer fmt.Println("Fim2")
	defer fmt.Println("Fim3")
	if numero > 5000 {
		fmt.Println("valor alto...")
		return 5000
	}

	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	//obterValorAprovado(6000)
	fmt.Println(obterValorAprovado(6000))
	// fmt.Println(obterValorAprovado(5000))

}
