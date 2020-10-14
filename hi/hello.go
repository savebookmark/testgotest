package main

import (
	"fmt"
	"math"
	"time"
)

func swap(x, y string) (string, string) {
	return y, x
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}
func hohoho() {
	var a int32
	var b int32
	var c int32
	a = 3
	b = 4
	c = a + b
	fmt.Println(c)
}
func say2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)

		fmt.Println(s)
	}

}
func main() {
	fmt.Println("hey")
	fmt.Println("Happy", math.Pi, "Day")
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	//////////////////////////////////////////////
	fmt.Println(".......................")
	//go say("world")
	//say("hello")
	hohoho()
	//say2("hahaha")
	//  c := [4]int{1, 2, 3, 4}
	//	c := []int{1, 2, 3, 4}
	c := make([]int, 4, 5)
	d := append(c, 5)

	d[0] = 33
	//슬라이스는 기본적으로 참조형이다 같은 주소를 참조하지만
	//cap 이 초과하면 새로운 메모리 공간을 가진다
	fmt.Printf("%d %d\n", c, d)
	fmt.Printf("%p %p\n", c, d)

}
