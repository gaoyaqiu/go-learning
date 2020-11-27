package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) func(op int) int {
	return func(n int) int {
		strt := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(strt).Seconds())
		return ret
	}
}


func slowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFn)
	t.Log(tsSF(10))
}

