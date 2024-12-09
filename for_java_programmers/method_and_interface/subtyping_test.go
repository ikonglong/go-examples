package method_interf

import (
	"testing"
)

func TestSubTyping(t *testing.T) {
	var p MySubType
	GetAndSet(&p, 10)
}

// 通过使用接口来实现方法调用的动态分派
func TestMethodCallDynamicallyDispatchedByUsingInterface(t *testing.T) {
	var v MyInterface

	v = new(MyType)
	v.Get() // Call the Get method for *MyType

	v = new(MySubType)
	v.Get() // Call the Get method for *MySubType
}

func TestUseAsSuperType(t *testing.T) {
	//var p MySubType
	// Compile error: cannot use 'p' (type MySubType) as the type MyType
	//var p2 MyType = p

	//p2 := MyType(p)
}
