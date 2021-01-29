package forexam

import "fmt"

func Guguan() {
	for dan := 1; dan <= 9; dan++ {
		fmt.Printf("%d단\n", dan)
		for j := 2; j <= 9; j++ {
			fmt.Printf("%d * %d = %dn", dan, j, dan*j)
		}
	}
	fmt.Println()
}
