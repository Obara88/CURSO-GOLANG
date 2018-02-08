package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("executou")
	fmt.Println("executou2")
	fmt.Println("executou3")

	ch <- 5
	ch <- 6

}

func main() {
	ch := make(chan int, 3)

	go rotina(ch)

	time.Sleep(time.Second * 1)

	fmt.Println(<-ch)

}
