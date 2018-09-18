package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var aa = 3
var ss = "kkk"

//使用var集中定义变量
var (
	aaa = 3
	sss = "kkk"
	bbb = true
)

//在func之外不能如下定义变量
//bb := true

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		//如果不指定类型，可以是int或者float
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	//常规写法
	/*const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)*/

	//iota特性1：<==使用iota数值自增==>
	const (
		cpp        = iota
		_
		java
		python
		golang
		javascript
	)

	//iota特性2：b, kb, mb, gb, tb, pb
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, aaa, sss, bbb)

	euler()
	triangle()
	consts()
	enums()
}
