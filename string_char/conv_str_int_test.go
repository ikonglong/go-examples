package string_char

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// https://yourbasic.org/golang/convert-int-to-string/

func TestIntToStr(t *testing.T) {
	// func name 'Itoa' means "integer to ASCII"
	s := strconv.Itoa(97)
	assert.Equal(t, "97", s)

	// Warning: In a plain conversion the value is interpreted as a Unicode code point,
	// and the resulting string will contain the character represented by that code point,
	// encoded in UTF-8.
	// s = string(97) // s == "a"
	// assert.Equal(t, "a", s)
}

func TestFormatIntToStrInGivenBase(t *testing.T) {
	var n int64 = 97
	s := strconv.FormatInt(n, 10) // s == "97" (decimal)
	assert.Equal(t, "97", s)

	s = strconv.FormatInt(n, 16) // s == "61" (hexadecimal)
	assert.Equal(t, "61", s)
}

func TestStrToInt(t *testing.T) {
	s := "97"
	n, _ := strconv.Atoi(s)
	assert.Equal(t, 97, n)

	// Use `strconv.ParseInt` to parse a decimal string (base 10) and check if it fits into an int64.
	// `strconv.ParseInt` returns (i int64, err error)
	n2, err := strconv.ParseInt(s, 10, 64)
	assert.Nil(t, err)
	assert.Equal(t, int64(97), n2)
	// The two numeric arguments represent a base (0, 2 to 36) and a bit size (0 to 64).
	//
	// If the first argument is 0, the base is implied by the string’s prefix:
	// base 16 for "0x", base 8 for "0", and base 10 otherwise.
	//
	// The second argument specifies the integer type that the result must fit into.
	// Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64.
}

func TestConvBetweenIntAndInt64(t *testing.T) {
	// The size of an int is implementation-specific, it’s either 32 or 64 bits, and hence
	// you won’t lose any information when converting from int to int64.
	var n int = 97
	m := int64(n)
	assert.Equal(t, int64(97), m)

	// However, when converting to a shorter integer type, the value is truncated
	// to fit in the result type's size.
	var v int64 = 2 << 32
	v2 := int(v)              // truncated on machines with 32-bit ints
	fmt.Printf("v2=%d\n", v2) // either 0 or 4,294,967,296
	assert.True(t, v2 == 0 || v2 == 4294967296)
}

func TestGeneralFormattingDataToStr(t *testing.T) {
	// The `fmt.Sprintf` function is a useful general tool for converting data to string
	s := fmt.Sprintf("%+8d", 97)
	// s == "     +97" (width 8, right justify, always show sign)
	assert.Equal(t, "     +97", s) // 包含数字和符号，宽度为 8
}
