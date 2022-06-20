package section28testing

import "fmt"

func benchmarkExam1() {
	fmt.Println(Greet("James"))
}

// Greet say hello to
func Greet(to string) string {
	return "Hello, " + to
}