package datastructure

import "fmt"

func DFSTest() {
	tree := Tree{}
	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Children); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Children[i].AddNode(val)
			val++
		}
	}
	/*
					1
				/	|	\
			2		3	 4
		 /\	 /\  /\
		5 6 7 8 9 10
	*/
	fmt.Println("\nDFS with recursive call")
	tree.DFS()
	fmt.Println("\nDFS with stack")
	tree.DFSWithStack()
}

func StackTest() {
	stack := NewStack()

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

func QueueTest() {
	queue := NewQueue()

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
	list := &LinkedList{}
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
