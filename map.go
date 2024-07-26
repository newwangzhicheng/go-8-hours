package main

import "fmt"

func main() {
	M1()
}

func M1() {
	m1 := make(map[string]bool, 1)
	m1["a"] = true
	m1["b"] = false
	fmt.Printf("%v, len=%d\n", m1, len(m1))
}
