package main

import "fmt"

func counter(start int) (func() int, func()) {
	// if the value gets mutated, the same is reflected in closure
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	// both ctr and incr have same reference to start
	// closures are created, but are not called
	return ctr, incr
}

func main() {
	// ctr, incr and ctr1, incr1 are different
	ctr, incr := counter(100)
	ctr1, incr1 := counter(100)
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())
	// incr by 1
	incr()
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1- ", ctr1())
	// incr1 by 2
	incr1()
	incr1()
	incr1()
	incr1()
	//porque o start incrementou????
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1- ", ctr1())
}

//http://keshavabharadwaj.com/2016/03/31/closure_golang/
//https://developer.mozilla.org/en-US/docs/Web/JavaScript/Closures
//https://stackoverflow.com/questions/111102/how-do-javascript-closures-work?rq=1

//https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
