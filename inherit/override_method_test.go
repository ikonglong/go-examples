package inherit

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

type iHandler interface {
	handle()
}

type handlerBase1 struct {
	iHandler
}

func (b *handlerBase1) handle() {
	// compile error: comparison of function handle != nil is always true
	//if b.iHandler != nil && b.iHandler.handle != nil {
	if b.iHandler != nil {
		b.iHandler.handle()
		return
	}
	fmt.Println("handlerBase1.handle(): ...")
}

type handlerBase2 struct {
	iHandler
}

type ConcreteHandler1 struct {
	*handlerBase1
}

type ConcreteHandler2 struct {
	*handlerBase2
}

func TestHandleIsNil(t *testing.T) {
	b1 := &handlerBase1{}
	assert.NotNil(t, b1.handle)

	c1 := &ConcreteHandler1{
		handlerBase1: b1,
	}
	b1.iHandler = c1
	// 注意，虽然从 ConcreteHandler1 源码看，它未重写 handle 方法，
	// 但它继承了 handlerBase1 实现的 handle 方法
	assert.NotNil(t, c1.handle)

	b2 := &handlerBase2{} // b2 not impl iHandler.handle()
	// panic: runtime error: invalid memory address or nil pointer dereference
	//assert.NotNil(t, b2.handle)

	c2 := &ConcreteHandler2{}
	// panic: runtime error: invalid memory address or nil pointer dereference
	//assert.Nil(t, c2.handle)

	b2.iHandler = c2
	// 神奇吧，handlerBase2 和 ConcreteHandler2 都未实现 handle 方法，在未将 c2 嵌入
	// b2 前 b2.handle 是空指针，嵌入后它就不是 nil 了
	assert.NotNil(t, b2.handle)

	c2.handlerBase2 = b2
	// 神奇吧，handlerBase2 和 ConcreteHandler2 都未实现 handle 方法，在未将 b2 嵌入
	// c2 前 c2.handle 是空指针，嵌入后它就不是 nil 了
	assert.NotNil(t, c2.handle)
}

func TestConcreteNotOverrideHandle(t *testing.T) {
	convey.Convey("Given only base type implements iHandler.handle(),"+
		" and base and concrete objects **refer to** each other", t, func() {
		b := &handlerBase1{}
		c := &ConcreteHandler1{
			handlerBase1: b,
		}
		b.iHandler = c
		convey.Convey("When call handle() on base or concrete object", func() {
			//b.handle()
			c.handle()
			convey.Convey("Then it results in err: stack overflow", func() {
			})
		})
	})
}

func TestBaseAndConcreteNotImplHandle(t *testing.T) {
	convey.Convey("Given both base and concrete not implement iHandler.handle()", t, func() {
		b := &handlerBase2{}
		c := &ConcreteHandler2{
			handlerBase2: b,
		}
		b.iHandler = c
		convey.Convey("When call handle() on base or concrete object", func() {
			b.handle()
			//c.handle()
			convey.Convey("Then it results in err: stack overflow", func() {
			})
		})
	})
}

func TestX(t *testing.T) {
	b := &handlerBase1{}
	c := &ConcreteHandler1{
		handlerBase1: b,
	}
	//b.iHandler = c
	b.handle()
	//c.handle()
	//fmt.Println(c.do1 == nil)
	//fmt.Println(c.handle == nil)
	var v iHandler = c
	_, ok := v.(interface {
		do2()
	})
	fmt.Println(ok)
}
