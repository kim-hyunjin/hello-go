package forexam

import "fmt"

// 별찍기
func Star() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			mt.Print("*")
		}
		mt.Println()
}