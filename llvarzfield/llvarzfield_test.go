package llvarzfield

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPack(t *testing.T) {
	assert := assert.New(t)
	field := New()
	field.Value = "1111111111111111D1212"
	result, _ := field.Pack()
	rs := hex.EncodeToString(result)
	assert.Equal("111111111111111111d12120", rs)
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	field := New()
	field.Value = "1111111111111111D1212"
	result, _ := field.Pack()
	field = New()
	field.Unpack(result)
	assert.Equal("1111111111111111d12120", field.Value)
}
