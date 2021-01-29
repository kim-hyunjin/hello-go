package datastructure

import "fmt"

func Slice1() {
	var a []int                         // 동적 배열 선언
	fmt.Printf("len(a) = %d\n", len(a)) // 0

	var b = []int{1, 2, 3, 4}
	fmt.Printf("len(b) = %d\n", len(b)) // 4

	var c = make([]int, 3, 8)           // 0 ~ 2번 인덱스는 0으로 채우고, capacity는 8
	fmt.Printf("len(c) = %d\n", len(c)) // 3
	fmt.Printf("cap(c) = %d\n", cap(c)) // 8
}

func Slice2() {
	var a []int
	fmt.Printf("len(a) = %d\n", len(a)) // 0
	fmt.Printf("cap(a) = %d\n", cap(a)) // 0
	a = append(a, 1)
	fmt.Printf("len(a) = %d\n", len(a)) // 0
	fmt.Printf("cap(a) = %d\n", cap(a)) // 0

}
