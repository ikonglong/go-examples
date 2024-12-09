package method_interf

import "fmt"

// This is not the same as a subclass in Java, but a form of delegation. When a method of an anonymous field is called,
// its receiver is the field, not the surrounding struct. In other words, methods on anonymous fields are not
// dynamically dispatched. When you want the equivalent of Java’s dynamic method lookup, use an interface.
type MySubType struct {
	MyType // an anonymous field of type MyType
	j      int
}

// he Get method was overridden, and the Set method was inherited.
func (p *MySubType) Get() int {
	p.j++
	fmt.Printf("[MySubType] i:%d, j:%d\n", p.i, p.j)
	return p.MyType.Get()
}

// The Set method is inherited from MyType, because methods associated with the anonymous field are promoted
// to become methods of the enclosing type.
