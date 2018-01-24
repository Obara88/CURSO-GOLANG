package main

import "fmt"

func main() {

	i := 3
	i += 3
	i -= 3
	i /= 3
	i *= 3
	i %= 3
	fmt.Println(i)

	x, y := 1, 2

	x, y = y, x

	fmt.Println(x, y)
}
