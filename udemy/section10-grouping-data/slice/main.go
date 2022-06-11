package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	var b [6]int 
	var c [5]int

	// different length -> different type
	fmt.Println("[5]int == [6]int ? ", reflect.TypeOf(a) == reflect.TypeOf(b))
	fmt.Println("[5]int == [5]int ? ", reflect.TypeOf(a) == reflect.TypeOf(c))

	x := []int{4, 5, 7, 8, 9}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:]) // slicing
	fmt.Println(x[1:3])

	x = append(x, 12)
	fmt.Println(x)

	y:= []int{123, 456, 789}
	x = append(x, y...)
	fmt.Println(x)
}