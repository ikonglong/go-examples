package struct_ex

import (
	"fmt"
	"testing"
)

type book struct {
	title   string
	price   int
	authors [1]person
	editor  person
}

type person struct {
	name string
}

func TestAssignValObjToVar(t *testing.T) {
	a := person{
		name: "x-man",
	}
	a2 := a
	fmt.Printf("a: %v, a2: %v\n", a, a2)
	fmt.Printf("&a: %p, &a2: %p\n", &a, &a2)
	a2.name = "batman"
	fmt.Printf("a: %v, a2: %v\n", a, a2)
}

func TestAssignValObjToObjField(t *testing.T) {
	p := person{
		name: "x-man",
	}
	var b book
	b.editor = p
	fmt.Printf("p: %v, b.editor: %v\n", p, b.editor)
	fmt.Printf("&p: %p, &b.editor: %p\n", &p, &(b.editor))
	p.name = "batman"
	fmt.Printf("p: %v, b.editor: %v\n", p, b.editor)
	fmt.Printf("&p: %p, &b.editor: %p\n", &p, &(b.editor))
}

func TestNewBook(t *testing.T) {
	theAuthor := person{
		name: "peter",
	}
	authors := [1]person{0: theAuthor}
	var price = 10
	var title = "go programming"
	fmt.Printf("&theAuthor: %p, &authors: %p, &(authors[0]): %p, &price: %p, &title: %p\n",
		&theAuthor, &authors, &(authors[0]), &price, &title)
	book := book{}
	book.title = title
	book.authors = authors
	book.price = price
	fmt.Printf("&book.authors: %p, &book.authors[0]: %p, &book.price: %p, &book.title: %p\n",
		&(book.authors), &(book.authors[0]), &(book.price), &(book.title))
}

func TestNewBook2(t *testing.T) {
	theAuthor := person{
		name: "peter",
	}
	authors := [1]person{0: theAuthor}
	var price = 10
	var title = "go programming"
	fmt.Printf("&theAuthor: %p, &authors: %p, &(authors[0]): %p, &price: %p, &title: %p\n",
		&theAuthor, &authors, &(authors[0]), &price, &title)
	book := book{
		title:   title,
		authors: authors,
		price:   price,
	}
	fmt.Printf("&book.authors: %p, &book.authors[0]: %p, &book.price: %p, &book.title: %p\n",
		&(book.authors), &(book.authors[0]), &(book.price), &(book.title))
	// Deep copy, not shallow copy
	fmt.Printf("&theAuthor.name: %p, &(book.authors[0].name): %p\n", &theAuthor.name, &(book.authors[0].name))
}
