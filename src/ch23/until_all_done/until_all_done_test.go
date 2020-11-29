package until_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Microsecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	// 防止协程泄露
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After: ", runtime.NumGoroutine())
}
