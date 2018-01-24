package main

import (
	"fmt"
)

func main() {
	fmt.Println("resultado", exec(mutiplicador, 1, 2))
	fmt.Println("resultado", exec(somar, 1, 2))
}

func mutiplicador(a, b int) int {
	return a * b
}

func somar(a, b int) int {
	return a + b
}

func exec(funcao func(x, y int) int, a, b int) int {
	return funcao(a, b)
}
