package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	// Zero value for slice is nil
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	//[2 4 6 8], len=4, cap=4
	printSlice(s1)

	s2 := make([]int, 16)
	//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s2)
	s3 := make([]int, 10, 32)
	//[0 0 0 0 0 0 0 0 0 0], len=10, cap=32
	printSlice(s3)

	copy(s2, s1)
	//[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s2)

	s2 = append(s2[:3], s2[4:]...)
	//[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=15, cap=16
	printSlice(s2)

	front := s2[0]
	s2 = s2[1:]
	//2 [4 6 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(front, s2)

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	//0 [4 6 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(tail, s2)
}
