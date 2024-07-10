package main

import "fmt"

var x, y int // 都是int
var (
	a bool = false
	b      = 2
) // 分解的写法用于声明全局变量

func main() {
	var x, y = false, 1
	fmt.Println(x, y)

	a, b := 1, 'b'
	fmt.Println(a, b)
}
