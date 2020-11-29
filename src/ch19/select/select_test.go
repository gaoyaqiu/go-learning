package my_select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 50)
	return "Done"
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1) // 更高效的做法
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Microsecond * 100):
		t.Error("time out")
	}
}
