package method_interf

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	var p = new(MyType)
	fmt.Printf("p.i:%d\n", p.Get())
}

func TestMyInt(t *testing.T) {
	var v MyInt
	v = v * v // The operators of the underlying type still apply.
	f(int(v)) // int(v) has no declared methods.
	//f(v)    // Compile error: cannot use 'v' (type MyInt) as the type int
}
