package string_char

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIntToStr(t *testing.T) {
	x := 123
	y := fmt.Sprintf("%d", x) // one option
	// func name 'Itoa' means "integer to ASCII"
	fmt.Println(y, strconv.Itoa(x)) // another option
}

func TestStrToNum(t *testing.T) {
	s := "123"
	x, _ := strconv.Atoi(s)             // x is an int
	y, _ := strconv.ParseInt(s, 10, 64) // base 10, up to 64 bits
	fmt.Println(x, y)
}

func TestFormatNumInDifferentBase(t *testing.T) {
	x := 123
	fmt.Println(strconv.FormatInt(int64(x), 2)) // output binary format
	fmt.Println(strconv.FormatUint(uint64(x), 2))
}
