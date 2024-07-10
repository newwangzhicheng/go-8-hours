package main

import "fmt"

const (
	TOP    = iota * 2 // 0
	RIGHT             // 2
	BOTTOM            // 4
	LEFT              // 6
)

func main() {
	fmt.Println("TOP", TOP)
}
