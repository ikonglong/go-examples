package error

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type domainError struct {
	msg   string
	cause error
}

func (e domainError) Error() string {
	return e.msg
}

func (e domainError) Unwrap() error {
	return e.cause
}

func operation() *domainError {
	return nil
}

func Test_ReturnNilDomainErrPtr_IsNil(t *testing.T) {
	f := func() *domainError {
		return nil
	}
	assert.True(t, f() == nil)
	fmt.Printf("f() == nil: %v\n", f() == nil)
}

func Test_ReturnNilErrInterface_IsNil(t *testing.T) {
	f := func() error {
		return nil
	}
	assert.True(t, f() == nil)
	fmt.Printf("f() == nil: %v\n", f() == nil)
}

func Test_ReturnNilDomainErrPtrByErrInterface_IsNil(t *testing.T) {
	f := func() error {
		var err *domainError
		return err
	}
	assert.True(t, f() != nil)
}

func Test_AfterAssignOpErrPtrNilToErrorIFace_IsNil(t *testing.T) {
	var err error
	fmt.Printf("before assign, (nil error) is nil: %v\n", err == nil)
	err = operation()
	fmt.Printf("after assign (nil *domainError) to (err error), err is nil: %v\n", err == nil)
}

type customErr interface {
	error
}

func Test_AfterAssignOpErrPtrNilToCustomErrorIFace_IsNil(t *testing.T) {
	var err customErr
	fmt.Printf("before assign, (nil customErr) is nil: %v\n", err == nil)
	err = operation()
	fmt.Printf("after assign (nil *domainError) to (err customErr), err is nil: %v\n", err == nil)
}

type serviceError struct {
	msg   string
	cause error
}

func (e serviceError) Error() string {
	return e.msg
}

func (e serviceError) Unwrap() error {
	return e.cause
}

func Test_stdlib_as_func(t *testing.T) {
	err := serviceError{
		msg: "service err ...",
		cause: domainError{
			msg:   "operation err ...",
			cause: errors.New("internal err ..."),
		},
	}

	// Compilation error:
	// second argument to errors.As must be a non-nil pointer to either a type that implements error,
	// or to any interface type
	//fmt.Printf("errors.As(err, serviceError{}): %v\n", errors.As(err, serviceError{}))

	// panic: errors: target must be a non-nil pointer
	//var sErr *serviceError
	//fmt.Printf("errors.As(err, (nil *serviceError)): %v\n", errors.As(err, sErr))

	var sErr serviceError
	fmt.Printf("sErr: %v\n", sErr)
	fmt.Printf("errors.As(err, (non-nil *serviceError)): %v\n", errors.As(err, &sErr))
	fmt.Printf("sErr: %v\n", sErr)

	fmt.Printf("========================\n")

	// 证明 errors.As(err, target) 也会在 error chain 中搜索目标错误，搜索到第一个就返回
	var opErr domainError
	fmt.Printf("opErr: %v\n", opErr)
	fmt.Printf("errors.As(err, (non-nil *domainError)): %v\n", errors.As(err, &opErr))
	fmt.Printf("opErr: %v\n", opErr)
}

func Test_stdlib_is_func(t *testing.T) {
	err := serviceError{
		msg: "service err ...",
		cause: domainError{
			msg:   "operation err ...",
			cause: errors.New("internal err ..."),
		},
	}

	fmt.Printf("errors.Is(err, serviceError{}): %v\n", errors.Is(err, serviceError{}))
	fmt.Printf("========================\n")

	var nilSErr *serviceError
	fmt.Printf("errors.Is(err, (nil *serviceError)): %v\n", errors.Is(err, nilSErr))
	fmt.Printf("========================\n")

	var sErr serviceError
	fmt.Printf("sErr: %v\n", sErr)
	fmt.Printf("errors.Is(err, (non-nil *serviceError)): %v\n", errors.Is(err, &sErr))
	fmt.Printf("sErr: %v\n", sErr)

	fmt.Printf("========================\n")

	// 证明 errors.Is(err, target) 也会在 error chain 中搜索目标错误，搜索到第一个就返回
	var opErr domainError
	fmt.Printf("opErr: %v\n", opErr)
	fmt.Printf("errors.Is(err, (non-nil *domainError)): %v\n", errors.Is(err, &opErr))
	fmt.Printf("opErr: %v\n", opErr)

	fmt.Printf("========================\n")

	eqOpErr := domainError{
		msg:   "operation err ...",
		cause: errors.New("internal err ..."),
	}
	// 虽然 eqOpErr 跟要查找的目标错误类型相同，值也想通，但仍然没有匹配上。为什么呢？
	// TODO 搞清楚比较逻辑
	fmt.Printf("errors.Is(err, eqOpErr): %v\n", errors.Is(err, eqOpErr))
	// 传入的查找目标值跟 error chain 中的匹配目标是同一个对象
	fmt.Printf("errors.Is(err, err.cause): %v\n", errors.Is(err, err.cause))
	target := err.cause.(domainError)
	// 传入的查找目标值跟 error chain 中的匹配目标是同一个对象
	fmt.Printf("errors.Is(err, err.cause): %v\n", errors.Is(err, target))
	// 为什么给 errors.Is(err, target) 传入 concrete type pointer 参数结果不符合预期呢？
	// 因为如果 target 是指针变量，interface value 的第二个字存储的是此指针变量的地址
	fmt.Printf("errors.Is(err, err.cause): %v\n", errors.Is(err, &target))

	// An error is considered to match a target if it is equal to that target or if
	// it implements a method Is(error) bool such that Is(target) returns true.
	// TODO 实现 Is 方法
}
