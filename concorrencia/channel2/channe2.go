package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	fmt.Println("entrei em sleep")
	time.Sleep(time.Second)
	fmt.Println("sai do sleep")
	c <- 2 * base
	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(99, c)
	fmt.Println("ja foi para essa instrução?")
	teste("primeiro teste")
	fmt.Println("Estou esperando os 2 canais")
	teste("segundo teste")
	a, b := <-c, <-c
	fmt.Println("a, b := <-c, <-c")
	fmt.Println(a, b)

	fmt.Println(<-c)

}

func teste(texto string) {
	fmt.Println(texto)
}

/*
ja foi para essa instrução?
entrei em sleep
primeiro teste
Estou esperando os 2 canais
segundo teste
sai do sleep
já li os 2 canais
4 6
8*/

/*
ja foi para essa instrução?
primeiro teste
Estou esperando os 2 canais
entrei em sleep
segundo teste
sai do sleep
a, b := <-c, <-c
4 6
8
*/
