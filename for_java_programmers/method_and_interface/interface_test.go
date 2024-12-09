package method_interf

import (
	"testing"
)

func TestInterface(t *testing.T) {
	var p MyType
	GetAndSet(&p, 10) // ok

	// Compile error: cannot use 'p' (type MyType) as the type MyInterface
	// Type does not implement 'MyInterface' as the 'Get' method has a pointer receiver
	//GetAndSet(p)
}
