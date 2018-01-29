package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBalisa()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBalisa() {
	fmt.Println("balisa...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBalisa()
}
