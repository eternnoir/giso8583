package tpdu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {
	assert := assert.New(t)
	tpdu, err := NewTpdu()
	if err != nil {
		assert.Fail(fmt.Sprint(err))
	}
	tpdu.Identifier = "60"
	tpdu.DestinationAddress = "0297"
	tpdu.OriginationAddress = "0000"
	result, err := tpdu.Pack()
	if err != nil {
		assert.Fail(fmt.Sprintf("TPDU pack fail. %s", err))
	}
	assert.Len(result, 5)
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	tpdu, err := NewTpdu()
	if err != nil {
		assert.Fail(fmt.Sprint(err))
	}
	tpdu.Identifier = "60"
	tpdu.DestinationAddress = "0297"
	tpdu.OriginationAddress = "0000"
	result, err := tpdu.Pack()
	tpdu, _ = NewTpdu()
	tpdu, err = tpdu.Unpack(result)
	if err != nil {
		assert.Fail(fmt.Sprintf("TPDU unpack fail. %s", err))
	}
	assert.Equal("0297", tpdu.DestinationAddress)
}
