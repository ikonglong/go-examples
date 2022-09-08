package num

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPrintFloat(t *testing.T) {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d, e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

// The smallest positive integer that cannot be exactly represented as a float32 is not large:
func TestSmallestPositiveIntThatCanNotBeExactlyRepresented(t *testing.T) {
	var f float32 = 16777216
	fmt.Println(f == f+1)
	assert.True(t, f == f+1)
}

func TestX(t *testing.T) {
	fmt.Printf("max float64: %16.f\n", math.MaxFloat64)

	// A float32 provides approximately six decimal digits of precision,
	// whereas a float64 provides about 15 digits.
	// 也就是说，float64 的 precision 为 16 位十进制数。
	var x float64 = 0
	x = 10567592720220710
	fmt.Printf("x = %f\n", x)
	x = 10567592720220711
	fmt.Printf("x = %f\n", x)
	x = 10567592720220713
	fmt.Printf("x = %f\n", x)
	x = 10567592720220715
	fmt.Printf("x = %f\n", x)

	x = 1056759272022070
	for d := 0; d <= 9; d++ {
		fmt.Printf("x = %f\n", x+float64(d))
	}
}
