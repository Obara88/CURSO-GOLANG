package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println(string(97)) //Essa conversao representa a seguinte pergunta, 97 corresponde com qual codigo da tabela ascii? Resposta: a = 97
	fmt.Println(string(98)) //Essa conversao representa a seguinte pergunta, 98 corresponde com qual codigo da tabela ascii? Resposta: b = 98

	a := 10 // isso é um int
	fmt.Println(a)

	b := float64(a) //isso é um float64
	fmt.Println(b)

	c := "texto" //isso é um texto
	fmt.Println(c)

	//como que eu concateno um numero + um texto?
	// converto o int para a string  ( que na verdade é uma array)
	fmt.Println(strconv.Itoa(10) + "texto")

	//como que eu print "Ben10" ?
	f := "Ben"
	g := 10
	fmt.Println(f + strconv.Itoa(g))

	//como que eu faço o inverso?
	//O que aconteceu aqui?
	// a funcao strconv.Atoi retonar 2 valores! o valor convertido e um erro
	// usamos o _ para que o go ignore a segunda variavel afinal o Go reclama se vc declarar variavel e nao utilizar
	d, _ := strconv.Atoi("10")
	fmt.Println(d)

	fmt.Println()

	//como que eu faço conversao de boolean?
	//o que convert "true" ou "True"
	// isso não funciona "trUE" ou "TrUe"

	e, _ := strconv.ParseBool("True")
	if e {
		fmt.Println("verdadeiro")
	} else {
		fmt.Println("false")
	}

}
