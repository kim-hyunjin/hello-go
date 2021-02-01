package main

import (
	"github.com/kim-hyunjin/hello-go/datastructure"
)

func main() {
	LinkedListTest()
}

func LinkedListTest() {
	list := &datastructure.LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNode()

	list.RemoveNode(list.Root)
	list.PrintNode()

	list.RemoveNode(list.Tail)
	list.PrintNode()

	list.RemoveNode(list.Root.Next)
	list.PrintNode()
	list.PrintReverse()
}
