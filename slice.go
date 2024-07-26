package main

import "fmt"

func main() {
	S1()
	S2()
	S3()
	S4()
	S5()
	S6()
}

func S1() {
	s1 := make([]int, 2, 4)
	fmt.Printf("%v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
}

func S2() {
	var s2 []int
	s2 = make([]int, 2)
	s2[0] = 100
	s2[1] = 101
	s2 = append(s2, 103)
	fmt.Printf("%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
}

func S3() {
	var s3 []int
	s3 = append(s3, 100)
	fmt.Printf("%v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
}

func S4() {
	var s4 []int
	s4 = append(s4, 100)
	ChangeSlice(s4)
	fmt.Printf("%v, len=%d, cap=%d\n", s4, len(s4), cap(s4))
}

func S5() {
	var s5 []int = []int{100, 101, 102}
	var s5Copy []int = make([]int, 2)
	copy(s5Copy, s5)
	fmt.Printf("%v, len=%d, cap=%d\n", s5Copy, len(s5Copy), cap(s5Copy))
}

func S6() {
	var s6 []int = []int{3, 2, 1}
	s6Slice := s6[:2]
	s6Slice[0] = -s6Slice[0]
	fmt.Printf("%v, len=%d, cap=%d\n", s6Slice, len(s6Slice), cap(s6Slice))
}

func ChangeSlice(s []int) {
	s[0] = s[0] + 1
}
