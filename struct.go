package main

import (
	"fmt"
	"math"
)

func main() {
	changeNameForT()
	calcDistance(
		Point{0, 3},
		Point{4, 0})
}

// 结构和方法
type T struct {
	name string
}

func (t *T) ChangeName() {
	t.name = "world"
}

func changeNameForT() {
	t := T{name: "hello"}
	t.ChangeName()
	fmt.Println(t.name) // 输出：hello world
}

// 方法值，其实就是把方法当作值传递
type Point struct {
	x float64
	y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func calcDistance(p Point, q Point) {
	distanceFromP := p.Distance
	fmt.Printf("distance from p is %f\n", distanceFromP(q))
}
