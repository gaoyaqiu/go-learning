package main

import (
	"fmt"
)

// 结合常量的iota 与<< 运算符实现计算机储存单位的枚举
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}
