package section14function

import (
	"fmt"
	"testing"
)

func TestReturningFunc(t *testing.T) {

	f := ReturnFunc()

	num := f()

	fmt.Printf("%T\n", f)
	fmt.Println(ReturnFunc()())

	if num != 451 {
		t.Error("not valid number")
	}
}