package main

import (
	"fmt"
	"os"
)

func main() {
	// 同包的函数
	fmt.Println("args:", os.Args)
	u := user()
	x := User()
	fmt.Println(u,x)
}
