# example01

## 定义变量

- 使用var默认值

```go
var a,b,c, bool
```

- 使用var指定初始值

```go
var s1,s2 string = "hello", "world"
```

- 可放在函数内，或者直接放包内
- 使用var集中定义变量
- 让编译器自动决定类型

```go
var a,b,i,s1,s2 = true,false,3,"hello","world"
```

- 使用`:=`定义变量

```go
a,b,i,s1,s2 := true,false,3,"hello","world"
```

- `:=`方式只能在函数内使用

# example02

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