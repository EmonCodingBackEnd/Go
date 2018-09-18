# example01

## 定义变量

- 使用var默认值

```go
var a,b,c bool
```

- 使用var指定初始值

```go
var s1,s2 string = "hello", "world"
```

- 可放在函数内，或者直接放包内
- 使用var集中定义变量

```go
var (
	aaa = 3
	sss = "kkk"
	bbb = true
)
```

- 让编译器自动决定类型

```go
var a,b,i,s1,s2 = true,false,3,"hello","world"
```

- 使用`:=`定义变量

```go
a,b,i,s1,s2 := true,false,3,"hello","world"
```

- `:=`方式只能在函数内使用



## 內建变量类型

- bool,string
- (u)int,(u)int8,(u)int16,(u)int32,(u)int64,(u)intptr
- byte,rune
- float32,float64,complex64,complex128

## 复数

- i =  $\sqrt{-1}$
- 复数： 3 + 4i

![复数](https://github.com/EmonCodingBackEnd/Go/blob/master/GoLearning/images/20180917073440.png)

- |3 + 4i| = $\sqrt{3^2 + 4^2}$ = 5
- $i^2$ = -1,$i^3$ = -i, $i^4$ = 1, ...

![复数2](https://github.com/EmonCodingBackEnd/Go/blob/master/GoLearning/images/20180917075857.png)

- $e^{iφ}$ = cosφ + isinφ
- |$e^{iφ}$| = $\sqrt{cos^2φ+sin^2φ}$ = 1
- $e^0$ = 1, $e^{i\frac{π}{2}}$ = i
- $e^{iπ}$ = -1, $e^{i\frac{3}{2}π}$ = -i, $e^{i2π}$ = 1

![复数3](https://github.com/EmonCodingBackEnd/Go/blob/master/GoLearning/images/20180917075117.png)

## 最美公式——欧拉公式

- $e^{iπ}$ + 1 = 0

![最美公式——欧拉公式](https://github.com/EmonCodingBackEnd/Go/blob/master/GoLearning/images/20180917080325.png)



## 强制类型转换

- 类型转换是强制的

```go
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
```

## 常量的定义

```go
func consts() {
	const (
		filename = "abc.txt"
		// 如果不指定类型，可以是int或者float
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
```

##  枚举类型

- 普通枚举类型
- 自增值枚举的类型

```go
func enums() {
    // 常规写法
	/*const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)*/
	
	// iota特性1：<==使用iota数值自增==>
	const (
		cpp        = iota
		_
		java
		python
		golang
		javascript
	)

	// iota特性2：b, kb, mb, gb, tb, pb
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
```

## 变量定义要点

- 变量类型写在变量名后面
- 编译器可以推测变量类型
- 没有char，只有rune
- 原生的支持复数类型

## 条件语句

- if的条件里面不需要括号

```go
// 常规写法
func normalIf() {
	const filename = "example/example01/branch.go"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}


```

- if的条件里面可以写变量

```go
// 特性写法
func featureIf() {
	const filename = "example/example01/branch.go"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
```

## switch

- switch会自动break，除非使用fallthrough

```go
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
```

- switch后可以没有表达式

```go
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
```

## for

- 常规写法

```go
// 常规写法
func normalFor() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

- 省略初始条件，相当于while

```go
//省略初始条件，相当于while
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
```

- 省略初始条件和递增条件，相当于while

```go
//省略初始条件和递增条件，相当于while
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
```

- 省略所有

```go
//省略所有条件，死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}
```

## 函数

定义：

```go
func eval(a, b int, op string) int
```

- 常规函数

```go
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
```

- 函数可以返回多个值

```go
func featureFunc(a, b int) (int, int) {
	return a / b, a % b
}
```

- 函数可以返回指定名称的值（仅用于非常简单的函数，禁止滥用）

```go
func featureFunc2(a, b int) (q, r int) {
	return a / b, a % b
}
```

- 多返回值时，一般最后一个参数是error

```go
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
// 示例调用
func main() {
	if result, err := featureFunc3(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
```

- 函数可以作为参考

```go
func highFeatureFunc(op func(int, int2 int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}
```

- 没有默认参数，可选参数等，只有一个可变参数

```go
func highFeatureFunc2(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
```

## 指针

Go语言只有值传递一种方式，不存在引用传递。

# example02

## 数组

- 数组是值类型
- `[10]int`和`[20]int`是不同类型
- 调用`func f(arr [10]int)`会**拷贝**数组
- 在go语言中一般不直接使用数组的