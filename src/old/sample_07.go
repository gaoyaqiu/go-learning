package main

import (
	"fmt"
)

// 指针学习： 操作符"&"取变量地址， 使用"*"通过指针间接访问目标对象
func main() {
	a := 1
	var p *int = &a
	// p指向的内存地址
	fmt.Println(p)
	// 取值时需要使用*号
	fmt.Println(*p)
}
