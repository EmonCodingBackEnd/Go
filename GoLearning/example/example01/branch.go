package main

import (
	"io/ioutil"
	"fmt"
)

//常规写法
func normalIf() {
	const filename = "example/example01/branch.go"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

//特性写法
func featureIf() {
	const filename = "example/example01/branch.go"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func normalSwitch(a, b int, op string) int {
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

func featureSwitch(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	normalIf()
	featureIf()
	fmt.Println(
		normalSwitch(1, 2, "+"),
		normalSwitch(1, 2, "-"),
		normalSwitch(1, 2, "*"),
		normalSwitch(1, 2, "/"),
		//normalSwitch(1, 2, "&"),
	)
	fmt.Println(
		featureSwitch(0),
		featureSwitch(59),
		featureSwitch(60),
		featureSwitch(82),
		featureSwitch(99),
		featureSwitch(100),
		//会中断switch的执行，包括其他正常的score值也不会被输出
		//featureSwitch(101),
		//featureSwitch(-3),
	)
}
