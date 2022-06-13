package section14function

import "fmt"

func FuncExpression() {
	f := func() {
		fmt.Println("my first func expression")
	}

	f()

	f2 := func(x int) {
		fmt.Println("the year big brother started watching: ", x)
	}

	f2(1984)
}