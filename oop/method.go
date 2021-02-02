package main

import "fmt"

type Jam struct {
}

func (b *Bread) string() string {
	return b.val
}
func (j *Jam) GetVal() string {
	return " + jam"
}

func (b *Bread) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}

func main() {
	bread := &Bread{val: "bread"}
	jam := &Jam{}

	bread.PutJam(jam)

	fmt.Println(bread.string())
}
