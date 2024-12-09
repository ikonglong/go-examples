package pointer

import (
	"fmt"
	"testing"
)

func TestBookPointer(t *testing.T) {
	var b *Book
	b = &Book{
		Title:  "Go Concurrency",
		Author: "X-man",
	}
	fmt.Printf("1) title: %s, author: %s\n", b.Title, b.Author)
	b = new(Book)
	b.Title = "Design Pattern"
	b.Author = "Martin fowler"
	fmt.Printf("2) title: %s, author: %s\n", b.Title, b.Author)
}
