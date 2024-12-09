package make_func

import (
	"fmt"
	"testing"
)

// Doc: https://yourbasic.org/golang/go-java-tutorial/#making-values

func TestMakeMap(t *testing.T) {
	var dict = make(map[string]int, 10)
	dict["a"] = 1
	fmt.Printf("val for key 'a' is %d", dict["a"])
}

func TestMakeSlice(t *testing.T) {
	var s = make([]int, 10) // both len and cap are 10
	s[8] = 60
	fmt.Printf("s[8]: %d\n", s[8])

	s = make([]int, 8, 10) // len is 8, cap is 10
	s[7] = 50
	fmt.Printf("s[7]: %d\n", s[7])
}

func TestMakeChannel(t *testing.T) {
	// todo
}
