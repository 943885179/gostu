package main

import (
	"fmt"
	"unsafe"
)

const (
	a1 = "abc"
	b1 = len(a1)
	c1 = unsafe.Sizeof(a1)
)

func main() { // 这个提示可以不用在意，因为同一个目录下面不能有个多 package main
	a, b := func1()
	fmt.Println(a, b)
	_, c := func1()
	fmt.Println(c)
	// 常量定义
	const t int = 10
	const (
		e1 = 0
		e2 = 1
	)
	println(a1, b1, c1, e1, e2, t)
	func2()
	func3(t)
	func3(11)
	func4(1)
	func4(2)
	func5(1)
	func5(0)
	func6(0)
}
func func1() (int, string) {
	a, b := 1, "123"
	return a, b
}
func func2() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
func func3(t int) {
	fmt.Println(t)
	if t == 10 {
		fmt.Println("If")
	} else {
		fmt.Println("else")
	}
}
func func4(x int) {
	switch x {
	case 1:
		fmt.Println("Switch 1")
	default:
		fmt.Println("Switch other")
	}
}
func func5(x int) {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case x = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
func func6(x int) {
	c3 := make(chan int)
	select {
	case c3 = <-c3:
		fmt.Printf("2131\n")
	default:
		fmt.Printf("no communicationaaa\n")
	}
}
func func7(t int) int {
	if t == 0 {
		return 0
	}
}
