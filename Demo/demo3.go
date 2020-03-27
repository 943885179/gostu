package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSequence())
	fmt.Println(getSequence())
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	// 等待匿名goroutine
	<-ch
	fmt.Println("all done")
	func_for()
}
func func_for() {
	for index := 0; index < 10; index++ {
		fmt.Print(index)
	}
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	/*
		for{
			//这个会无限循环
		}*/
}
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
