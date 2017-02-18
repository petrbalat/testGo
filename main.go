package main

import (
	"fmt"
	"time"
)

func add(a int, b int) int  {
	return a + b
}

func main() {
	fmt.Println("Hello, Petr")
	fmt.Println("Time is ", time.Now())
	fmt.Println(add(4, 5))
}
