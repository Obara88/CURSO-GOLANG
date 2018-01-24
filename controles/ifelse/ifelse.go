package main

import (
	"fmt"
)

func main() {
	if 1 == 1 {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	texto := "teste"

	if texto == "a" {
		fmt.Println(texto)
	} else if texto == "b" {
		fmt.Println(texto)
	} else if texto == "c" {
		fmt.Println(texto)
	} else if texto == "e" {
		fmt.Println(texto)
	} else {
		fmt.Println(texto)
	}

	//if com declaração de variavel
	if teste := valor(); teste {
		fmt.Println(teste)
	} else {
		fmt.Println(teste)
	}
}

func valor() bool {
	return true
}
