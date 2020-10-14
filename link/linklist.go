package main

import "fmt"

//Node .
type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root
	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
		fmt.Println(tail)
	}
	PrintNode(root)
	RemoveNode(root)
	PrintNode(root)
}

// AddNode ..
func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

// RemoveNode ..
func RemoveNode(prenode *Node) *Node {
	prenode.next = prenode.next.next
	return prenode
}

// PrintNode ..
func PrintNode(root *Node) {
	pnode := root
	// fmt.Println(pnode)
	for pnode.next != nil {
		fmt.Printf("| %d next=%d |", pnode.val, pnode.next)
		pnode = pnode.next
	}
	// 마지막 tail은 넥스트가 없으므로
	fmt.Printf("| %d %d |\n", pnode.val, pnode.next)
}
