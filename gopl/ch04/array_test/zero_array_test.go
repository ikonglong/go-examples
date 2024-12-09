package array_test

import (
	"fmt"
	"testing"
)

func zero_v1(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero_v2(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func TestZeroV1(t *testing.T) {
	a := [...]int{1, 2, 3}
	fmt.Println(a)
}

func TestZeroV2(t *testing.T) {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
}
