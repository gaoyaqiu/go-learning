package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n must be greater than 2")
var GreaterThanHundredError = errors.New("n must be less than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		//return nil, errors.New("n should be in [2,100]")
		return nil, GreaterThanHundredError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	var list []int
	list, err := GetFibonacci(5)

	if err == LessThanTwoError {
		t.Error("Need a larger number")
		return
	}
	t.Log(list)
}

/**
异常处理最佳实践，通常会反着写，减少起嵌套、快速失败
*/
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}

	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}


func TestGetFibonacci2(t *testing.T) {
	GetFibonacci2("10")
}

