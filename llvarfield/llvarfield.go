package llvarfield

import (
	"encoding/hex"
	util "github.com/eternnoir/giso8583/util"
	"strconv"
)

type LLVARField struct {
	Length int
	Value  string
}

func New() *LLVARField {
	return &LLVARField{}
}

func (field *LLVARField) Pack() ([]byte, error) {
	field.Length = len(field.Value)
	lengthHeader := strconv.Itoa(field.Length / 2)
	rawHexData := util.PadLeft(lengthHeader, "0", 2) + field.Value
	return hex.DecodeString(rawHexData)
}

func (field *LLVARField) Unpack(byteary []byte) (*LLVARField, error) {
	rawHex := hex.EncodeToString(byteary)
	withoutLenHex := rawHex[2:len(rawHex)]
	field.Value = withoutLenHex
	field.Length = len(field.Value)
	return field, nil
}
