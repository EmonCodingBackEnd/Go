package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

//常规写法
func normalFor() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//省略初始条件，相当于while
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

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

//省略所有条件，死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	normalFor()
	fmt.Printf("%s,%s,%s,%q",
		convertToBin(5),
		convertToBin(13),
		convertToBin(72387885),
		convertToBin(0),
	)
	printFile("example/example01/loop.go")
	forever()
}
