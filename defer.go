package main

import "fmt"

func P(num int) {
	fmt.Println(num)
}

func main() {
	defer P(1)
	defer P(2)
	defer P(3)
	defer P(4)
}
