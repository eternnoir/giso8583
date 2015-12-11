package bitmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackUnpack(t *testing.T) {
	assert := assert.New(t)
	bitmap := New(64)
	bitmap.SetBit(5, 1)
	bitmap.SetBit(10, 1)
	result, _ := bitmap.Pack()
	assert.Equal(len(result), 8)
	bitmap = New(64)
	bitmap.Unpack(result)
	assert.Equal(int(bitmap.GetBit(1)), 0)
	assert.Equal(int(bitmap.GetBit(5)), 1)
	assert.Equal(int(bitmap.GetBit(10)), 1)
}
