package main

import "fmt"

func updateSlice(s []int) {
	//由于传递的是原数组的slice，会改变arr的数据
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//arr[2:6] =  [2 3 4 5]
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:]
	updateSlice(s1)
	//[100 3 4 5 6 7]
	fmt.Println(s1)

	s2 := arr[:]
	updateSlice(s2)
	//[100 1 100 3 4 5 6 7]
	fmt.Println(s2)
	//[100 1 100 3 4 5 6 7]
	fmt.Println(arr)

	s2 = s2[:5]
	s2 = s2[3:]
	//[3 4]
	fmt.Println(s2)

	updateSlice(s2)
	//[100 4]
	fmt.Println(s2)
	//[100 1 100 100 4 5 6 7]
	fmt.Println(arr)

	//数据复原
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[2:6]
	s2 = s1[3:5]
	//s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s2=[5 6], len(s2)=2, cap(s2)=3
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

}
