package string

import "testing"

func TestString(t *testing.T)  {
	var s string
	t.Log(s) // 初始化默认值""
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' // string是不可变的byte slice
	s = "\xE4\xB8\xA5" // 可以存储二进制数据
	t.Log(s)
	t.Log(len(s)) // 求出来的是里面的byte数，不一定是字符数
	s = "中"
	t.Log(len(s)) // 求出来的是里面的byte
	
	c := []rune(s)
	t.Log(len(c))
	
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s)
}

func TestStringToRune(t *testing.T)  {
	s := "中华人民共和国"
	for _, c:= range s{
		t.Logf("%[1]c %[1]x", c)
	}
}

