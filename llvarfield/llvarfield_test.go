package llvarfield

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
	assert.Equal(rs, "028014")
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
