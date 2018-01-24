package main

import (
	"fmt"
)

//assim que se declara funcao com retorno o retorno é logo apos os parentes
func somar(a int, b int) int {
	return a + b
}

func imprimir(a int) {
	fmt.Println(a)
	d, e, f := teste("", "", "")

	fmt.Println(d, e, f)

	fmt.Println(teste("a", "b", "c"))
}

// rodar o codigo no windows
// C:\Users\kenji\Desktop\CURSO-GOLANG\fundamentos\funcao> go run main.go funcoes.go

func teste(batata, banana, maca string) (string, string, string) {
	return batata, banana, maca
}
