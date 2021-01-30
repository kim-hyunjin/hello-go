package datastructure

import "fmt"

type Node struct {
	next *Node
	prev *Node // 더블 링크드 리스트
	val  int
}
type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(target *Node) {
	if target == l.root {
		l.root = l.root.next
		l.root.prev = nil
		target = nil
		return
	}

	// 지울 노드의 이전 노드 찾기 => 더블링크드리스트로 만들어서 prev를 찾는 부분이 필요없어짐.
	// prev := l.root
	// for prev.next != node {
	// 	prev = prev.next
	// }

	if target == l.tail {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		target.prev.next = target.next
		target.next.prev = target.prev
	}
	target = nil
}

func (l *LinkedList) PrintNode() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverse() {
	node := l.tail

	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func LinkedListTest() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNode()

	list.RemoveNode(list.root)
	list.PrintNode()

	list.RemoveNode(list.tail)
	list.PrintNode()

	list.RemoveNode(list.root.next)
	list.PrintNode()
	list.PrintReverse()
}
