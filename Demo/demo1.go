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
	// twoSum([]int{1,2,3},3)
}
func twoSum(nums []int, target int) []int {
        for a,i := range nums{
            for b,j := range nums{
                if a+b == target{
                    return []int{a,b}
                }
            }
        }
}
