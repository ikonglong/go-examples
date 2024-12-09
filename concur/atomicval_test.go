package concur

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
	"time"
)

// atomic.Value 包装值变量
func TestAtomicValueWrapValVar(t *testing.T) {
	book1 := book{
		title: "OOP Design",
	}
	fmt.Printf("&book1: %p\n", &book1)

	var v atomic.Value
	v.Store(book1)

	book1.title = book1.title + " V1"
	fmt.Printf("book1: %v, v.Load(): %v\n", book1, v.Load().(book))

	// compile error: Cannot assign to v.Load().(book).title
	//v.Load().(book).title = "OOP Design V1"

	innerVal := v.Load()
	fmt.Printf("&book1: %p, &(v.Load()): %p\n", &book1, &innerVal)
}

// atomic.Value 包装引用变量
func TestAtomicValueWrapRefVar(t *testing.T) {
	book := book{
		title: "OOP Design",
	}
	var v atomic.Value
	v.Store(&book)
	book.title = book.title + " V1"
	fmt.Printf("book: %v, &book: %p, v.Load(): %p\n", book, &book, v.Load())
}

// 并发访问 atomic.Value 类型的结构体数据成员
func TestConcurAccessAtomicValueMember(t *testing.T) {
	aBook := book{
		title: "OOP Design",
	}
	aBook.soldFlag.Store(false)
	assert.False(t, aBook.isSold())

	aBook.onBookSold()
	time.Sleep(50 * time.Millisecond)

	assert.True(t, aBook.isSold())
}

type book struct {
	title    string
	soldFlag atomic.Value
}

func (b *book) onBookSold() {
	go func() {
		fmt.Printf("receieved book-sold event\n")
		b.soldFlag.Store(true)
	}()
}

func (b *book) isSold() bool {
	return b.soldFlag.Load().(bool)
}
