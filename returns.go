package main

import "fmt"

func main() {
	name, age := getProfile2()
	fmt.Printf("name %s, age %d\n", name, age)
}

func getProfile() (string, int) {
	return "Alice", 25
}

func getProfile2() (name string, age int) {
	name = "Jay"
	age = 52
	return
}
