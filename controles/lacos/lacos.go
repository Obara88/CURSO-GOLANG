package main

import (
	"fmt"
	"time"
)

func main() {

	//forma de usar loop

	//while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//for classico
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//for infinito
	for {
		fmt.Println("1")
		time.Sleep(time.Second)
	}

	forrange1()
}

func forrange1() {
	//inicialização de array e o compilador que ira definir o tamanho para vc
	numeros := [...]int{9, 98, 2, 2486, 74, 56723, 41567, 5646, 13}

	//é um foreach da vida
	for indice, conteudo := range numeros {
		fmt.Println(indice, conteudo)
	}

	fmt.Println("")
	//só pegando o conteudo da array, omitindo o primeiro parametro ( indice)
	for _, conteudo := range numeros {
		fmt.Println(conteudo)
	}

	fmt.Println("")

	//só pegando o indice da array, pra que? nao sei...mas da
	for i := range numeros {
		fmt.Println(i)
	}

}
