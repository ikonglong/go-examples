package method_interf

import (
	"fmt"
)

type MyType struct {
	i int
}

func (p *MyType) Get() int {
	fmt.Printf("[MyType] i:%d\n", p.i)
	return p.i
}

type MyInt int

func (p MyInt) Get() int {
	//return p // Compile error: cannot use 'p' (type MyInt) as the type int
	return int(p) // The conversion is required.
}

func f(i int) {}
