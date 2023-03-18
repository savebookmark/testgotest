package main

import (

	// "example/test/calcal"
	"example/test/calcal"
	"example/test/channeltest"
	"example/test/dataStruct"

	// "example/test/docktyping"
	// "example/test/doublelinklist"
	// "example/test/hi"
	// "example/test/interfacetest"
	// "example/test/link"
	"example/test/mutexlock"
	// "example/test/selecttest"
	"fmt"
	"sync"
	"time"

	"rsc.io/quote/v4"
)

func main() {
	// queuestack()
	// queuestack2()

	treeS()
	fmt.Println()
	fmt.Println("---------------")
	bitree()
	fmt.Println("---------------")
	heapadd()
	fmt.Println("---------------")
	heapmin()
	fmt.Println("---------------")
	// lock()
	// fmt.Println("---------------")
	// channelCarmake()

	fmt.Println(quote.Hello())

	// baseball.Bsmain()
	calcal.Calmain()
	// docktyping.Dockmain()
	// doublelinklist.Doublemain()
	// hi.Himain()
	// interfacetest.Interfacemain()
	// link.Linkmain()
	// selecttest.Selectmain()

}
func queuestack() {
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
	fmt.Println("//////////////////////////////////")

}
func queuestack2() {
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

	fmt.Println()
	fmt.Println("NewQueue")

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d ->", val)
	}
	fmt.Println()
	fmt.Println("//////////////////////////////////")
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
func heapmin() {

	h := &dataStruct.MinHeap{}

	// [-1, 3, -1, 5, 4], 2번째 큰값
	nums := []int{-1, 3, -1, 5, 4}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())

	//Input: [2, 4, -2, -3, 8], 1
	//Output: 8
	h = &dataStruct.MinHeap{}

	nums = []int{2, 4, -2, -3, 8}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 1 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())

	//Input: [-5, -3, 1], 3
	//Output: -5
	h = &dataStruct.MinHeap{}

	nums = []int{-5, -3, 1}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 3 {
			h.Pop()
		}
	}
	fmt.Println(h.Pop())
}
func lock() {
	for i := 0; i < 20; i++ {
		mutexlock.Accounts = append(mutexlock.Accounts, &mutexlock.Account{Balance: 1000, Mutex: &sync.Mutex{}})
	}
	mutexlock.GlobalLock = &sync.Mutex{}

	mutexlock.PrintTotalBalance()

	for i := 0; i < 10; i++ {
		go mutexlock.GoTransfer()
	}

	for {
		mutexlock.PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}

// 채널 사용
func channelCarmake() {
	chan1 := make(chan channeltest.Car)
	chan2 := make(chan channeltest.Car)
	chan3 := make(chan channeltest.Car)

	go channeltest.StartWork(chan1)
	go channeltest.MakeTire(chan1, chan2)
	go channeltest.MakeEngine(chan2, chan3)

	for {
		result := <-chan3
		fmt.Println(result.Val)
	}
}
