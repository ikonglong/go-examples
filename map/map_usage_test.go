package mapusage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateNilMap(t *testing.T) {
	var m map[string]any
	for k, v := range m {
		println("%s: %v", k, v)
	}
}

func TestGetFromNilMap(t *testing.T) {
	var m map[string]any
	assert.Nil(t, m)
	v, found := m["k1"]
	assert.True(t, v == nil)
	assert.False(t, found)
}

func TestPutToNilMap(t *testing.T) {
	var m map[string]any
	assert.Nil(t, m)
	// panic: assignment to entry in nil map [recovered]
	//     panic: assignment to entry in nil map
	m["k1"] = "v1"
	v, found := m["k1"]
	assert.True(t, v != nil)
	assert.True(t, found)
}

func TestMapContain(t *testing.T) {
	m := map[string]interface{}{}
	v, found := m["a"]
	assert.Nil(t, v)
	assert.False(t, found)

	v2 := m["a"]
	assert.Nil(t, v2)
}
