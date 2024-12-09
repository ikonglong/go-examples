package array_test

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x,\n%x,\n%t,\n%T,\n", c1, c2, c1 == c2, c1)
}
