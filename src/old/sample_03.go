package main

import (
	"fmt"
	"strconv"
)

// 定义别名组
type (
	byte int8
	rune int32
	ByteSize int64
)

// 只有全局变量的声明才可以使用var()的方式来简写，不能用在局部变量中
var (
	// 使用常规方式
	aaa = "hello"
	// 使用并行方式及其类型推断
	bbb, ccc = 1, 2
)

func main() {
	// 没有赋值时默认为0
	var i int
	// 没有赋值时默认false
	var b bool
	// 没有赋值时默认为空字符串
	var s string
	// 定义1个长度的数组
	var a [1]byte

	// 省略var关键字，由编译器自己推断数据类型
	c := 11111

	// 定义多个参数并赋值
	var a1, a2, a3 = 1, 2, 3

	// go中不存在隐式转换， 所有类型转换必须显示声明
	// 而且转换只能发生在两种相互兼容的类型之间
	var ff float32 = 100.1
	fmt.Println(ff)
	fi := int(ff)
	fmt.Println(fi)
	// float不能转换成bool类型， 两种类型不同
	//fb := bool(ff)

	var g1 int = 45
	// int 转换成 string
	g2 := strconv.Itoa(g1)

	var gstr string = "123"
	// 将string 转换为 int
	g3, err := strconv.Atoi(gstr)
	if err != nil {
		fmt.Printf("gstr %s is not an integer - exiting with error\n", gstr)
		return
	}

	fmt.Println(g1)
	fmt.Println(g2)
	fmt.Println(g3)

	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

}
