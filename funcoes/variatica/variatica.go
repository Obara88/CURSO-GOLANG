package main

import (
	"fmt"
)

func main() {
	//variatica1()
	variaticaSlice()
}

func variatica1() {
	fmt.Println("media", media(12, 3.21, 31.3, 123.21, 3.0, 1.3, 1.23, 213.0, 23.0))
}

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))

}

func variaticaSlice() {
	aprovados := []string{"Joao", "Paulo", "Ana"} //isso é um slice
	// aprovados := [...]string{"Joao", "Paulo", "Ana"} //isso é um array

	//quando vc faz isso "aprovados..." ele quebra em varias string
	imprimirAprovados(aprovados...)

}

func imprimirAprovados(aprovados ...string) { //trato isso "aprovados ..." como um array
	for indice, valor := range aprovados {
		fmt.Println(indice, valor)
	}
}
