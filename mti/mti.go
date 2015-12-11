package mti

import (
	"encoding/hex"
	"errors"
	"fmt"
)

const (
	REQ_SALE = "0200"
	RES_SALE = "0210"
)

type Mti struct {
	Value string
}

func (mti *Mti) Pack() ([]byte, error) {
	if mti.Value == "" {
		return nil, errors.New(fmt.Sprintf("MTI must have value."))
	}
	ret, err := hex.DecodeString(mti.Value)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("MTI encode error. %s", err))
	}
	return ret, nil
}

func (mti *Mti) Unpack(byteary []byte) (*Mti, error) {
	mti.Value = hex.EncodeToString(byteary)
	return mti, nil
}
