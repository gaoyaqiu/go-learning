package main

import (
	"fmt"
)

// 常量组定义，如果不提供初始值，则表示使用上行的表达式
// 如果不想在外部被调用可以在常量名前加_ 或者小写字母
const (
	MAX_COUNT = 1
	MIN_COUNT
	_MAX_NUMBER = 3
)

// iota 是常量计数器， 从0开始，组中每定义一个常量，自动递增1
// 每遇到一个const关键字， iota就会重置为0
const (
	NUMBER1 = iota
	NUMBER2 = iota
)

const (
	NUMBER3 = iota
	NUMBER4
)

func main() {
	fmt.Println(MAX_COUNT)
	fmt.Println(MIN_COUNT)
	fmt.Println(NUMBER1)
	fmt.Println(NUMBER2)
	fmt.Println(NUMBER3)
	fmt.Println(NUMBER4)
}
