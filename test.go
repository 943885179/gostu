package main

import "fmt"

func main() {
	var a string = "weiqiao"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var e bool
	fmt.Println(e)
	fmt.Printf("%v %q\n", a, b)
	f := true
	fmt.Println(f)
	x := twoSum([]int{1, 2, 3}, 3)
	fmt.Println(x)
}
func twoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums {
			if a+b==target && a!=b {
				return []int{i, j}
			}
		}
	}
	return nil
}
