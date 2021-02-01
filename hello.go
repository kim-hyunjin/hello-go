package main

import (
	"fmt"

	"github.com/kim-hyunjin/hello-go/datastructure"
)

func main() {
	stackTest()
	queueTest()
}

func stackTest() {
	stack := datastructure.NewStack()

	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}

	fmt.Println("New Stack")

	for !stack.Empty() {
		val := stack.Pop()
		fmt.Println(val)
	}
	// stack := []int{}
	// for i := 1; i <= 5; i++ {
	// 	stack = append(stack, i)
	// }

	// fmt.Println(stack)

	// for len(stack) > 0 {
	// 	var last int
	// 	last, stack = stack[len(stack)-1], stack[:len(stack)-1]
	// 	fmt.Println(last)
	// }
}

func queueTest() {
	queue := datastructure.NewQueue()

	for i := 1; i <= 5; i++ {
		queue.Push(i)
	}

	fmt.Println("New Queue")
	for !queue.Empty() {
		val := queue.Pop()
		fmt.Println(val)
	}

	// queue := []int{}
	// for i := 1; i <= 5; i++ {
	// 	queue = append(queue, i)
	// }

	// fmt.Println(queue)

	// for len(queue) > 0 {
	// 	var front int
	// 	front, queue = queue[0], queue[1:]
	// 	fmt.Println(front)
	// }
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
