package main

import (
	"go-8-hours/module1"
	_ "go-8-hours/module2"
	. "go-8-hours/module3"
)

func main() {
	module1.PrintA()
	PrintC()
}
