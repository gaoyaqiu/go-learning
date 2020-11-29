package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	// defer 函数始终会被执行
	//defer func() {
	//	fmt.Println("Finally!")
	//}()

	defer func() {
		// 常见的错误恢复
		if err := recover(); err != nil {
			// 小心僵尸进程
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
}
