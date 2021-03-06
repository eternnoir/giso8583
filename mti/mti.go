package mti

import (
	"encoding/hex"
	"errors"
	"fmt"
)

type Mti struct {
	Value string
}

func New() *Mti {
	return new(Mti)
}

func (mti *Mti) Pack() ([]byte, error) {
	if mti.Value == "" {
		return nil, errors.New("MTI must have value.")
	}
	ret, err := hex.DecodeString(mti.Value)
	if err != nil {
		return nil, fmt.Errorf("MTI encode error. %s", err)
	}
	return ret, nil
}

func (mti *Mti) Unpack(byteary []byte) (*Mti, error) {
	mti.Value = hex.EncodeToString(byteary)
	return mti, nil
}
