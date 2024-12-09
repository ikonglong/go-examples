package string_char

import (
	"fmt"
	"testing"
)

func TestChangeString(t *testing.T) {
	s := "hello"
	// Compile error: Cannot assign to s[4]
	//s[4] = 'x'
	fmt.Println(s)
}
