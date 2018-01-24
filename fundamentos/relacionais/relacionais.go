package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("String:", "Banana" == "Banana")
	fmt.Println("String:", "banana" == "Banana")

	d1 := time.Unix(0, 0)
	fmt.Println(d1)
}
