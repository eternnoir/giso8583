package bitmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackUnpack(t *testing.T) {
	assert := assert.New(t)
	bitmap := New(63)
	bitmap.Set(5)
	bitmap.Set(10)
	result, _ := bitmap.Pack()
	assert.Equal(len(result), 8)
	bitmap = New(63)
	bitmap.Unpack(result)
	f1, _ := bitmap.Get(1)
	f5, _ := bitmap.Get(5)
	f10, _ := bitmap.Get(10)
	assert.Equal(f1, false)
	assert.Equal(f5, true)
	assert.Equal(f10, true)
}
