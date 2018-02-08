package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second * 1)
	c <- 1
	fmt.Println("Só depois que foi lido")
	teste()
}

func main() {
	c := make(chan int)

	go rotina(c)

	fmt.Println(<-c) //operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c)
	fmt.Println("Fim") //deadlock
}

func teste() {
	for index := 0; index < 999999999; index++ {

	}
	fmt.Println("999...")
}
