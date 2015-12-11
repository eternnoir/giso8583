package lllvarfield

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {
	assert := assert.New(t)
	field := New()
	field.Value = "8014"
	result, _ := field.Pack()
	rs := hex.EncodeToString(result)
	assert.Equal(rs, "00028014")
}

func TestPacklong(t *testing.T) {
	value := "11111111111111111111"
	assert := assert.New(t)
	field := New()
	field.Value = value
	result, _ := field.Pack()
	rs := hex.EncodeToString(result)
	assert.Equal(rs, "0010"+value)
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	field := New()
	field.Value = "8014"
	result, _ := field.Pack()
	field = New()
	rs, _ := field.Unpack(result)
	assert.Equal(rs.Value, "8014")
}

func Testunacklong(t *testing.T) {
	value := "11111111111111111111"
	assert := assert.New(t)
	field := New()
	field.Value = value
	result, _ := field.Pack()
	field = New()
	rs, _ := field.Unpack(result)
	assert.Equal(rs.Value, value)
}
