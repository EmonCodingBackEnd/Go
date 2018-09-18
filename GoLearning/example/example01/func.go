package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//常规函数
func normalFunc(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operation:" + op)
	}
	return result
}

//函数可以返回多个值
func featureFunc(a, b int) (int, int) {
	return a / b, a % b
}

//函数可以返回指定名称的值
func featureFunc2(a, b int) (q, r int) {
	return a / b, a % b
}

//多返回值时，一般最后一个参数是error
func featureFunc3(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := featureFunc2(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation:" + op)
	}
}

//函数可以作为参数
func highFeatureFunc(op func(int, int2 int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

//没有默认参数，可选参数等，只有一个可变参数
func highFeatureFunc2(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(normalFunc(3, 4, "*"))
	q, r := featureFunc2(13, 3)
	fmt.Println(q, r)

	if result, err := featureFunc3(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(highFeatureFunc(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(highFeatureFunc2(1, 2, 3, 4, 5))
}
