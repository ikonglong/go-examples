package array_test

import (
	"fmt"
	"testing"
)

func TestAccessByIndex(t *testing.T) {
	var a [3]int
	fmt.Println("First elem: ", a[0])
	fmt.Println("Last elem: ", a[len(a)-1])

	var b [3]int = [3]int{1, 2, 3}
	fmt.Println("b[2]: ", b[2])
}

func TestIterate(t *testing.T) {
	var a [3]int = [3]int{1, 2, 3}

	// 输出索引和元素
	for i, v := range a {
		fmt.Printf("%d->%d\n", i, v)
	}

	// 仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

func TestOutputTypeOfArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	// 数组长度是数组类型的一部分
	fmt.Printf("Type of array a: %T\n", a)

	// var a1 = [3]int{1, 2, 3}
	// a1 = [4]int{1, 2, 3, 4} // 编译错误：Cannot use '[4]int{1, 2, 3, 4}' (type [4]int) as the type [3]int
}
