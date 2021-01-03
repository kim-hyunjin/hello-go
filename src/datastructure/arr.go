package datastructure

import "fmt"

func Arr() {
	// A:= [10]int{1,2,3,4,5,6,7,8,9,10}
	B:= [10]int {}
	for i:=0; i < len(B); i++ { // len : 배열의 길이를 반환하는 내장함수
		B[i] = i*i
	}
	fmt.Println(B)
}

func Arr2() {
	s:="hello 월드" // go 문자열은 utf-8 사용. s는 byte배열임.

	fmt.Printf("len(s) = %d\n", len(s))

	for i:=0; i<len(s); i++ {
		fmt.Println(string(s[i])) // byte배열을 순회하면 1바이트씩 읽기때문에 3바이트 한글은 깨짐
	} 
}

func Arr3() {
	s := "hello 월드"
	s2:=[]rune(s) // rune은 utf-8을 담는 데이터 타입
	fmt.Printf("len(s2) = %d\n", len(s2))

	for i:=0; i<len(s2); i++ {
		fmt.Println(string(s2[i]))
	} 
}