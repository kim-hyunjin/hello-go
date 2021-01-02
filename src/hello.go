package main

import (
	"fmt"
	_ "forexam"
	"functest"
)

func main() {
	fmt.Println("hello, go")
	fmt.Printf("총합: %d\n", functest.Sum2(10, 0))
	fmt.Printf("피보나치: %d\n", functest.Fibonacci(10))
}
