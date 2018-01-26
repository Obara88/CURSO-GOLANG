package main

import (
	"fmt"
)

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

//o resultado foi 20 e 10 porque a funcao tem "ciencia" do escopo que ela foi definida e naquele escopo x é 10, por isso que ao executar o imprimeX o valor impresso foi 10

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//quando se executou intSeq, foi criado outro contexto e assim terá outro valores
	newInts := intSeq()
	fmt.Println(newInts())

	incriment, get := Mycounter(10)
	for index := 0; index < 5; index++ {
		incriment()
	}
	fmt.Println("The final value is", get())

}

//o i permanece com o seu contexto
func intSeq() func() int {
	i := 0

	var funcao = func() int {
		i++
		return i
	}
	return funcao
}

type incFuncType func()
type getFuncType func() int

//funcao que retorno 2 funcoes, uma void e outra com retorno de inteiro
func Mycounter(initialCount int) (incFuncType, getFuncType) {
	theCount := initialCount

	increment := func() {
		theCount++
		fmt.Println("The count is", theCount)
	}
	get := func() int {
		fmt.Println("The Count is", theCount)
		return theCount
	}

	return increment, get
}
