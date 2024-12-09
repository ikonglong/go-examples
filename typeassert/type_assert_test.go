package typeassert

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// reference docs:
// https://go.dev/tour/methods/15
// https://yourbasic.org/golang/type-assertion-switch/

func TestAssertTypeOfNilInterfaceVal(t *testing.T) {
	var v interface{} = nil
	assert.Nil(t, v)
	_, ok := v.(interface{})
	fmt.Printf("v: %T(%#v)\n", v, v)
	assert.False(t, ok)
}

func TestAssertTypeOfNilMap(t *testing.T) {
	var m map[string]interface{}
	fmt.Printf("m: %T(%#v)\n", m, m)
	assert.Nil(t, m)

	// compile error:
	// Invalid type assertion: m.(map[string]interface{}) (non-interface type map[string]interface{} on the left)
	// _, ok := m.(map[string]interface{})

	var v interface{} = m
	fmt.Printf("v: %T(%#v)\n", v, v)
	assert.Nil(t, v)

	concreteV, ok := v.(map[string]interface{})
	fmt.Printf("concreteV: %T(%#v)\n", concreteV, concreteV)
	assert.True(t, ok)
	assert.Nil(t, concreteV)
}

func TestWhenASwitchCaseIsUsedForMultiTypes(t *testing.T) {
	var s *string = nil
	assert.Nil(t, s)

	// A type switch matches the dynamic type of the interface value 'x'.
	// The dynamic type is matched against the types in 'switch' cases.
	// If a short variable assignment of the form 'v := x.(type)' is used
	// as the switch guard and a switch case is used for a single type only,
	// 'v' will have the type specified in the matching switch case.

	// compile error:
	// Invalid type switch guard: v := s.(type) (non-interface type *string on the left)
	// switch v := s.(type) {
	var x any = s
	switch v := x.(type) {
	case *string:
		assert.Nil(t, v)
		assert.True(t, v == nil)
	}

	switch v := x.(type) {
	case *string, string: // a switch case is used for multi types
		assert.Nil(t, v) // ok，因为 Nil(...) 方法内部判断了接口对象中的 value 是否为 nil
		// 注意，这里的结果跟前一个 type switch 相反
		assert.False(t, v == nil)
	}

	var i *int = nil
	assert.Nil(t, i)
	x = i
	switch v := x.(type) {
	case *int:
		assert.Nil(t, v)
		assert.True(t, v == nil)
	}

	switch v := x.(type) {
	case *int, *int64:
		assert.Nil(t, v)
		assert.False(t, v == nil)
	}

	switch v := x.(type) {
	case *int, int:
		assert.Nil(t, v)
		assert.False(t, v == nil)
	}

	switch v := x.(type) {
	case int:
		assert.Nil(t, v)
		// assert.False(t, v == nil) // compile error: Cannot convert 'nil' to type 'int'
	}
	switch v := x.(type) {
	case int, int64:
		assert.Nil(t, v)
		assert.False(t, v == nil) // no compile error
	}
}
