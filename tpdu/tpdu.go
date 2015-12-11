package tpdu

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func New() (*Tpdu, error) {
	tpdu := Tpdu{}
	tpdu.IdentifierLength = 1
	tpdu.DestinationAddressLength = 2
	tpdu.OriginationAddressLength = 2
	return &tpdu, nil
}

type Tpdu struct {
	Identifier               string
	IdentifierLength         int
	DestinationAddressLength int
	DestinationAddress       string
	OriginationAddressLength int
	OriginationAddress       string
}

func (tpdu *Tpdu) Pack() ([]byte, error) {
	if tpdu.DestinationAddress == "" {
		return nil, errors.New("DestinationAddress can not be empty.")
	}
	if tpdu.OriginationAddress == "" {
		return nil, errors.New("OriginationAddress can not be empty.")
	}

	// Decode Identifier to byte array
	id, err := hex.DecodeString(tpdu.Identifier)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("TPDU Identifier decode error. %s", err))
	}

	// Decode Destination Address to byte array.
	da, err := hex.DecodeString(tpdu.DestinationAddress)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("TPDU DestinationAddress decode error. %s", err))
	}

	// Decode Origination Address to byte array.
	oa, err := hex.DecodeString(tpdu.OriginationAddress)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("TPDU OriginationAddress decode error. %s", err))
	}
	ret := []byte{}
	ret = append(ret, id...)
	ret = append(ret, da...)
	ret = append(ret, oa...)
	return ret, nil
}

func (tpdu *Tpdu) Unpack(byteary []byte) (*Tpdu, error) {
	hexString := hex.EncodeToString(byteary)
	offset := 0
	tpdu.Identifier = hexString[offset : offset+tpdu.IdentifierLength*2]
	offset = offset + tpdu.IdentifierLength*2
	tpdu.DestinationAddress = hexString[offset : offset+tpdu.DestinationAddressLength*2]
	offset = offset + tpdu.DestinationAddressLength*2
	tpdu.OriginationAddress = hexString[offset : offset+tpdu.OriginationAddressLength*2]
	return tpdu, nil
}
