package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	//funcao1()
	//pilha1()
	x, y := retornoNomeado(1, 2)
	fmt.Println(x, y)
}

func funcao1() {
	f1()
	f2("", "")

	//atribuição do retorno na var a
	a := f3("")
	fmt.Println(a)

	//atribuição do retorno na var b e c
	b, c := f5("", "")
	fmt.Println(b, c)

	//segundo parametro ignorado
	d, _ := f5("", "")
	fmt.Println(d)

	//primeiro parametro ignorado
	_, e := f5("", "")
	fmt.Println(e)

	//chamada do f5, o uso do retorno nao é obrigatorio
	f5("", "")

	//multipla atribuição
	g, h := f3("a"), f3("b")
	fmt.Println(g, h)

}

func f1() {
	//sem parametro e sem retorno
}

func f2(p1, p2 string) {
	// escrita de parametro reduziada
}

func f3(p1 string) string {
	return p1
}

func f5(p1, p2 string) (string, string) {
	return p1, p2
}

func pilha1() { p1() }
func p1()     { p2() }
func p2()     { p3() }
func p3()     { p4() }
func p4() {
	debug.PrintStack()
}

func retornoNomeado(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return //returno limpo

}
