package collection_test

import (
	"fmt"
	"github.com/ikonglong/go-examples/for_java_programmers/collection"
	"testing"
)

func TestStack(t *testing.T) {
	var s collection.Stack
	s.Push("World!")
	s.Push("Hello ")
	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}
	fmt.Println()
}
