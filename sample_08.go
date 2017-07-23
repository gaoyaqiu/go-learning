package main

import (
	"fmt"
)

// 控制语句学习
func main() {
	// 判断语句if
	a := 10
	if a := 5; a > 1 {
		// 里面的a是局部变量，不会影响到外部的a
		fmt.Println(a)
	}
	// 这里a是10
	fmt.Println(a)

	// 循环语句 for， go中只有for一个循环语句关键字， 但支持3种形式
	b := 1
	for {
		b++
		if b > 3 {
			break
		}
		fmt.Println(b)
	}
	fmt.Println("第一种形式over")

	c := 1
	for c <= 3 {
		c++
		fmt.Println(c)
	}
	fmt.Println("第二种形式over")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("第三种形式over")

	// 选择语句 switch， 不需要写break，一旦条件符合自动终止，如果希望执行下一个case，需使用
	// fallthrough 语句
	j := 1
	switch j {
	case 0:
		fmt.Println("j=0")
	case 1:
		fmt.Println("j=1")
	default:
		fmt.Println("None")
	}
	fmt.Println("switch 经典用法 over")

	switch {
	case j >= 0:
		fmt.Println("j>0")
		// 如果想继续执行下一个case，使用fallthrough
		fallthrough
	case j >= 1:
		fmt.Println("j>=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("switch 第二种用法 over")

	// switch带初始化语句用法
	switch g := 1; {
	case g >= 0:
		fmt.Println("g>0")
		// 如果想继续执行下一个case，使用fallthrough
		fallthrough
	case g >= 1:
		fmt.Println("g>=1")
	default:
		fmt.Println("None")
	}
	fmt.Println("switch 第三种用法 over")

	// 跳转语句goto（调整执行位置），break，continue
	// break与continue配合标签可用于多层循环的跳出
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				// 如果break改为goto或者continue，就会变成死循环
				break LABEL1
			}
		}
	}
	fmt.Println("ok")
}
