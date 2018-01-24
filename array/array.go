package main

import (
	"fmt"
)

func main() {
	//array1()
	//slice1()
	//sliceMake()
	//sliceCopy()
	//map1()
	map2()

}

func array1() {
	var notas [3]float64
	notas[0], notas[1], notas[2] = 1, 2, 3.6
	fmt.Println(notas)

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
		fmt.Println(notas[i])
	}

	media := total / float64(3)
	fmt.Println(media)
}

func slice1() {
	// define um array com 10 elementos cujos tipos são bytes
	var ar = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//define um slite! basicamente é um ponteiro para manipular array! ou algo do tipo...
	slice := ar[:7] // [0:n-1]

	fmt.Println(slice)

	arrayasdsa := [5]int{0, 1, 2, 3, 4}
	slice1 := arrayasdsa[:4]
	fmt.Println(arrayasdsa, slice1)

}

func sliceMake() {
	//o make criar um array interno e referencia atraves de um slice, de tamanho X
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// S foi atribuido um novo slice que tem 10 posições e a array interna que o slice aponta tem um tamanho de 20
	s = make([]int, 10, 20) //a leitura é : crie pra mim um slite que aponta ate a 10 posição de uma array com tamanho 20
	fmt.Println(s, len(s), cap(s))

	//O Make é a mesma coisa que declarar um arra e depois fazer o slice para ela
	b := [20]int{} //array de 20 posições
	c := b[:10]    //slice aponta para esta array porem referencia somente ate a posição 10

	c = append(c, 1, 2, 3)
	fmt.Println(c, len(c), cap(c))

}

func sliceCopy() {
	array := [3]int{1, 2, 3}
	var slice1 []int
	slice1 = append(slice1, 4, 5, 6) //o append redimensiona a array interna q o slice aponta caso o tamanho exceda, mesmo quando nao aponte para nenhuma array
	fmt.Println(array, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // o copy ele vai copiar a menor, e pronto. o slice1 tinha 3 posições e o slice2 tinha 2, ele só copiou até dois
	fmt.Println(slice2)

}

func map1() {
	// maps deve ser inicializado!
	aprovados := make(map[int]string)

	aprovados[123211] = "Maria"
	aprovados[123212] = "Pedro"
	aprovados[123213] = "João"
	aprovados[123214] = "Paulo"

	//declara e inicializar map

	mapinicializado := map[string]int{
		"a": 23123,
		"b": 23123,
		"c": 23123,
		"d": 23123,
	}

	fmt.Println(mapinicializado)

	for chave, valor := range aprovados {
		fmt.Println("chave:", chave, "valor:", valor)
	}

	fmt.Println(aprovados[123211])
	delete(aprovados, 123211)
	fmt.Println("deletado", aprovados[45564564564])
}

func map2() {
	usuario := map[string]map[string]int{
		"a": {
			"alba": 10,
			"abu":  20,
		},
		"b": {
			"barnei":  30,
			"bigorna": 40,
		},
		"c": {
			"claus":   50,
			"cripton": 60,
		},
	}

	//fmt.Println(usuario)

	for indice := range usuario {
		outroMap := usuario[indice]
		for indice := range outroMap {
			fmt.Println(indice, outroMap[indice])
		}
	}

}
