package anfield

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {
	assert := assert.New(t)
	field := ANField{}
	field.Value = "AAA"
	result, _ := field.Pack()
	assert.Equal(len(result), 3)
	assert.Equal(field.Length, 3)
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	field := &ANField{}
	field.Value = "AAA"
	result, _ := field.Pack()
	field = &ANField{}
	field, _ = field.Unpacke(result)
	assert.Equal(field.Value, "AAA")
}
