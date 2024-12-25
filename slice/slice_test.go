package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type book struct {
	title string
}

func revise(books []book) []book {
	fmt.Printf("[revise func] addr of books: %p, books as reference: %p\n", &books, books)
	books[0].title = "Go programming"
	return books
}

func TestMakeSlice(t *testing.T) {
	// make(sliceType, len):
	// 底层数组大小为 len。slice 容量和长度都为 len，start 为 0
	slice1 := make([]int, 2)
	fmt.Printf("%v\n", slice1)
	// make(sliceType, len, cap):
	// 底层数组大小为 cap。slice 容量为 cap，长度为 len，start 为 0
	slice2 := make([]int, 2, 2)
	fmt.Printf("%v\n", slice2)
	slice3 := make([]int, 0, 2)
	fmt.Printf("%v\n", slice3)
}

func TestPassSlice(t *testing.T) {
	books := []book{
		book{},
	}
	fmt.Printf("[before revise] addr of books: %p, books as reference: %p\n", &books, books)
	fmt.Printf("[before revise] addr of books[0]: %p, books[0]: %v\n", &(books[0]), books[0])
	returnedBooks := revise(books)
	fmt.Printf("[after revise] addr of returnedBooks: %p, retunedBooks as reference: %p\n", &returnedBooks,
		returnedBooks)
	fmt.Printf("[after revise] addr of books[0]: %p, books[0]: %v\n", &(books[0]), books[0])
}

func TestPanicWhenAccessSliceOfWhichTheLenIs0(t *testing.T) {
	// panic: runtime error: index out of range [0] with length 0 [recovered]
	//    panic: runtime error: index out of range [0] with length 0
	strs := make([]string, 0, 3)
	// 因为 strs 长度是 0，所以下标 0 越界
	strs[0] = "a"
}

func TestOkWhenAccessSliceOfWhichTheLenIs0(t *testing.T) {
	strs := make([]string, 0, 3)
	// 如果 len 为 0，append 会修改 len
	strs = append(strs, "a")
	assert.Equal(t, "a", strs[0])
}

func TestAppendToSlice(t *testing.T) {
	strs := make([]string, 3, 3)
	strs = append(strs, "d")
	// 注意，append 不是从头开始追加，而是追加在尾部。它将新元素追加在 len 位置，只要 len 小于等于 capacity
	assert.Equal(t, "", strs[0])
	assert.Equal(t, "d", strs[3])
}

func TestCustomSliceType(t *testing.T) {
	type names []string
	l := make(names, 0)
	l = append(l, "a")
	assert.Equal(t, "a", l[0])

	l2 := make([]names, 0)
	l2 = append(l2, names{"a"})
	assert.Equal(t, names{"a"}, l2[0])
}
