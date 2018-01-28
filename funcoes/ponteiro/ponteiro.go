package main

import (
	"fmt"
)

func main() {
	n := 1

	inc1(n) //por valor
	fmt.Println(n)

	inc2(&n) //por referencia
	fmt.Println(n)
}

func inc1(n int) {
	n++
}

//ponteiro para inteiro
func inc2(n *int) {
	*n++
}
