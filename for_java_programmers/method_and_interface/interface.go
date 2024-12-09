package method_interf

type MyInterface interface {
	Get() int
	Set(i int)
}

// Since MyType already has a Get method, we can make MyType satisfy the interface by adding
func (p *MyType) Set(i int) {
	p.i = i
}

// Now any function which takes MyInterface as a parameter will accept a variable of type *MyType.
func GetAndSet(x MyInterface, i int) {
	x.Get()
	x.Set(i)
}
