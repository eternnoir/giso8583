package mti

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {
	assert := assert.New(t)
	mti := Mti{}
	mti.Value = REQ_SALE
	result, err := mti.Pack()
	if err != nil {
		assert.Fail(fmt.Sprintf("MTI Pack error. %s", err))
	}
	assert.Equal(len(result), 2)
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	mti := Mti{}
	mti.Value = REQ_SALE
	result, err := mti.Pack()
	if err != nil {
		assert.Fail(fmt.Sprintf("MTI Pack error. %s", err))
	}
	mti = Mti{}
	unpackResult, err := mti.Unpack(result)
	assert.Equal(unpackResult.Value, REQ_SALE)

}
