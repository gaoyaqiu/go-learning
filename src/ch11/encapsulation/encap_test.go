package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e) // 变成指针类型
	t.Logf("e2 is %T", e2)
}

/**
	这种方式会复制实例中的内容（不推荐使用）
 */
//func (e Employee) String() string {
//	fmt.Printf("String() Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}
/**
使用这种方式并没有复制产生，所有的实例数据是存在同一块位置（内存地址）上的
 */
func (e *Employee) String() string {
	fmt.Printf("String() Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T)  {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("TestStructOperations() Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
