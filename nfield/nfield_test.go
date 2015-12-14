package nfield

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {
	assert := assert.New(t)
	field := New()
	field.Value = "000000"
	result, _ := field.Pack()
	assert.Equal(len(result), 3)
	assert.Equal(field.Length, 6)
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	field := New()
	field.Value = "000000"
	result, _ := field.Pack()
	field = New()
	field.Unpack(result)
	assert.Equal(field.Value, "000000")
	assert.Equal(field.Length, 6)
}
