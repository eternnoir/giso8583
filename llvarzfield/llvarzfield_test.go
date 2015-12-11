package llvarzfield

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {
	assert := assert.New(t)
	field := New()
	field.Value = "1111111111111111D1212"
	result, _ := field.Pack()
	rs := hex.EncodeToString(result)
	assert.Equal(rs, "111111111111111111D12120")
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	field := New()
	field.Value = "1111111111111111D1212"
	result, _ := field.Pack()
	field = New()
	rs, _ := field.Unpack(result)
	assert.Equal(rs.Value, "1111111111111111D12120")
}
