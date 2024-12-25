package sort_

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortCharsInStr(t *testing.T) {
	s := "dbca5"
	bytes := []byte(s)
	// panic: reflect: call of Swapper on string Value
	// sort.Slice(s, func(i, j int) bool { ... })
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	assert.Equal(t, "5abcd", string(bytes))
}
