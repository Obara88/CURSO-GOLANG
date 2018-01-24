package main

import (
	"fmt"
	"time"
)

func main() {

	switch1()
	switch2()
	switch3()
}

func switch1() {
	var teste = 10
	var nota = int(teste)
	texto := ""

	switch nota {
	case 10:
		texto += "passei por aqui mas continuei "
		fallthrough
	case 9:
		texto += "9"
	case 8:
		texto += "8"
	case 7:
		texto += "7"
	case 6:
		texto += "6"
	default:
		texto += "default"
	}

	fmt.Println(texto)

}

func switch2() {

	agora := time.Now()

	switch {
	case agora.Hour() <= 12:
		fmt.Println("ainda é cedo")
	case agora.Hour() > 12 && agora.Hour() <= 18:
		fmt.Println("tarde!")
	case agora.Hour() >= 18:
		fmt.Println("noite")
	default:
		fmt.Println("...")
	}
}
