package main

import (
	"fmt"
)

func main() {

	a := 1
	b := 2
	//declaração implicita de C ser um ponteiro para o tipo INT
	//c := &a
	//var c = &a

	fmt.Println("endereço de a", &a, "valor de a", a)
	fmt.Println("endereço de b", &b, "valor de b", b)

	c := &a
	fmt.Println("c aponta para a :", c, "endereco do a", &a, "conteudo do c ", *c, "valor de a", a, "endereco do c", &c)

	c = &b
	fmt.Println("c aponta para b :", c, "endereco do b", &b, "conteudo do c ", *c, "valor de b", b, "endereco do c", &c)

	c = &a
	fmt.Println("c aponta para c :", c, "conteudo do c", &b, "conteudo do c ", *c, "valor de a", b, "endereco do c", &c)

}
