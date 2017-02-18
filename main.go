package main

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	return a + b
}

func swap(par1, par2 string) (string, string) {
	return par2, par1
}

var pom = false

func main() {
	fmt.Println("Hello, Petr")
	fmt.Println("Time is ", time.Now())
	fmt.Println(add(4, 5))

	println("----------------")
	var a, b = swap("a", "b")
	println(a, b)

	println("----------------")

	if pom {
		println("true")
	} else {
		println("false")
	}

	println("----------------")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
