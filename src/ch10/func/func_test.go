package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int  {
	return func(n int) int {
		strt := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(strt).Seconds())
		return ret
	}
}

func slowFn(op int) int  {
	time.Sleep(time.Second * 1)
	return op
}


func TestFn(t *testing.T) {
	//a, b := returnMultiValues()
	//t.Log(a, b)
	// 获取一个返回值
	c, _ := returnMultiValues()
	t.Log(c)
	tsSF := timeSpent(slowFn)
	t.Log(tsSF(10))
}
