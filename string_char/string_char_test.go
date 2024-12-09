package string_char

import (
	"errors"
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestCountBytesAndChars(t *testing.T) {
	s := "Hello, 世界"
	fmt.Println("byte count:", len(s))
	fmt.Println("char/rune count:", utf8.RuneCountInString(s))

	c := 0
	//for _, _ = range s {
	for range s {
		c++
	}
	fmt.Println("char/rune count:", c)
}

func Test1TraverseCharsInStr(t *testing.T) {
	s := "Hello, 世界"
	for i := 0; i < len(s); {
		// Each call to DecodeRuneInString returns r, the rune itself,
		// and size, the number of bytes occupied by the UTF-8 encoding of r.
		r, size := utf8.DecodeRuneInString(s[i:])
		// %c	the character represented by the corresponding Unicode code point
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

func Test2TraverseCharsInStr(t *testing.T) {
	s := "Hello, 世界"
	for i, r := range s {
		// %d	base 10
		// %q	a single-quoted character literal safely escaped with Go syntax.
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

func TestConvIntToRuneValue(t *testing.T) {
	// Converting an integer value to a string interprets the integer as a rune value,
	// and yields the UTF-8 representation of that rune.
	fmt.Println(string(int32(65)))     // "A", not 65
	fmt.Println(string(int32(0x4eac))) // "京"

	// If the rune is invalid, the replacement character is substituted
	fmt.Println(string(int32(1234567))) // "�"
}

func TestDiffBetweenTwoChars(t *testing.T) {
	s := "cd"
	a := 'a'
	for _, r := range s {
		fmt.Printf("diff bwtween %c and %c: %d\n", a, r, r-a)
	}
}

func TestFindChar(t *testing.T) {
	s := "a你b好xg"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == 'x' {
			fmt.Println("find char x, so break")
			break
		}
		i += size
	}
}

func TestPrintChar(t *testing.T) {
	ch := 'A'
	fmt.Printf("%8b %d %c \n", ch, ch, ch)
	errors.Is(nil, nil)
}
