package main

import "fmt"

// 常量定义
const PI = 3.14

// 全局变量的声明与赋值
var name = "go"

// 一般类型声明
type i int

// 结构声明
type go1 struct{}

// 接口声明
type go2 interface{}

func main() {
	i := 1
	fmt.Println(PI)
	fmt.Println(name)
	fmt.Println(i)
}
