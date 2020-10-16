package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	stack := []int{}

	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}

	stack2 := dataStruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	fmt.Println("NewStack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)
	}

	queue2 := dataStruct.NewQueue()
	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}

	fmt.Println("NewQueue")

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d ->", val)
	}

	treeS()
	bitree()
	heapadd()
}
func treeS() {
	tree := dataStruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}
	tree.DFS1()
	fmt.Println()
	//hahahah branch m
	tree.DFS2()
}
func bitree() {
	tree := dataStruct.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()
	fmt.Println()

	if found, cnt := tree.Search(6); found {
		fmt.Println("found 6 cnt:", cnt)
	} else {
		fmt.Println("not found 6 cnt:", cnt)
	}

	if found, cnt := tree.Search(11); found {
		fmt.Println("found 11 cnt:", cnt)
	} else {
		fmt.Println("not found 11 cnt:", cnt)
	}
}
func heapadd() {
	h := &dataStruct.Heap{}

	h.Push(2)
	h.Push(6)
	h.Push(9)
	h.Push(6)
	h.Push(7)
	h.Push(8)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
