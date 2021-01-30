package datastructure

import "fmt"

type Node struct {
	next *Node
	val  int
}

func LinkedListTest() {
	var root *Node
	var tail *Node
	root = &Node{val: 0}
	tail = root

	// for i := 1; i < 10; i++ {
	// 	AddNode(root, i)
	// }

	// PrintNode(root)

	for i := 1; i < 10; i++ {
		tail = AddNode2(tail, i)
	}
	PrintNode(root)

	root, tail = RemoveNode(root.next, root, tail)
	PrintNode(root)

	root, tail = RemoveNode(root, root, tail)
	PrintNode(root)

	root, tail = RemoveNode(tail, root, tail)
	PrintNode(root)
}

func AddNode(root *Node, val int) {
	var tail *Node
	tail = root
	for tail.next != nil {
		tail = tail.next
	}

	node := &Node{val: val}
	tail.next = node
}

func AddNode2(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	// 지울 노드의 이전 노드 찾기
	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = node.next
	}

	return root, tail
}

func PrintNode(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
