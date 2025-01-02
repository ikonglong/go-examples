package hash

import (
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha256(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("a"))
	hv := h.Sum(nil)
	assert.True(t, len(hv) > 0)
	b := make([]byte, 0, 256)
	hv2 := h.Sum(b)
	fmt.Printf("hv : %x\nb  : %x\nhv2: %x\n", hv, b, hv2)
	assert.Equal(t, hv, hv2)
}

func TestReuseHashObj(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("a"))
	fmt.Printf("\"a\" hash: %x\n", h.Sum(nil))

	h.Reset()
	h.Write([]byte("b"))
	h2 := sha256.New()
	h2.Write([]byte("b"))
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Printf("%x\n", h2.Sum(nil))
	assert.Equal(t, fmt.Sprintf("%x", h.Sum(nil)), fmt.Sprintf("%x", h2.Sum(nil)))
}
