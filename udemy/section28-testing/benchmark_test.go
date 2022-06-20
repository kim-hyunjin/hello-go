package section28testing

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	benchmarkExam1()
	s := Greet("James")
	if s != "Hello, James" {
		t.Error("got", s, "expected", "Hello, James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Hello, James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}