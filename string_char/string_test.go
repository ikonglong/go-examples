package string_char

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// string 是只读的，不可修改
func TestChangeString(t *testing.T) {
	s := "abcd"
	// Compile error: Cannot assign to s[3]
	// s[3] = 'x'
	assert.Equal(t, "abcd", s)
}

func TestAppendIntAsCharsToStr(t *testing.T) {
	s := "a"
	sBytes := []byte(s)
	n := 1

	sBytes2 := append(sBytes, byte(n))
	assert.Equal(t, "a\u0001", string(sBytes2)) // why???

	sBytes3 := append(sBytes, byte('0'+n)) // '0'+n 将 n 转换为字符，但只适用于单位数。
	assert.Equal(t, "a1", string(sBytes3)) // why???

	n2 := 10
	sBytes4 := append(sBytes, byte('0'+n2)) // '0'+n 只适用于单位数。对于多位数，结果不是期望的。
	assert.Equal(t, "a:", string(sBytes4))  // why???

	// 最推荐的方法
	sBytes5 := append(sBytes, strconv.Itoa(n2)...)
	assert.Equal(t, "a10", string(sBytes5))
}
