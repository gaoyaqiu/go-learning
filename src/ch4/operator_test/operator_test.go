package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)


func TestBitClear(t *testing.T) {
	a := 7 // 0111
	//a = a &^ Readable // 清除只读权限
	//a = a &^ Writable // 清除写权限
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}


func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b) // false
	//t.Log(a == c)
	t.Log(a == d) // true
}
