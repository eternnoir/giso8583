package lllvarfield

import (
	"encoding/hex"
	"strconv"

	util "github.com/eternnoir/giso8583/util"
)

type LLLVARField struct {
	Length int
	Value  string
}

func New() *LLLVARField {
	return new(LLLVARField)
}

func (field *LLLVARField) Pack() ([]byte, error) {
	field.Length = len(field.Value)
	lengthHeader := strconv.Itoa(field.Length / 2)
	rawHexData := util.PadLeft(lengthHeader, "0", 4) + field.Value
	return hex.DecodeString(rawHexData)
}

func (field *LLLVARField) Unpack(byteary []byte) (*LLLVARField, error) {
	rawHex := hex.EncodeToString(byteary)
	withoutLenHex := rawHex[4:len(rawHex)]
	field.Value = withoutLenHex
	field.Length = len(field.Value)
	return field, nil
}

func (field *LLLVARField) GetValue() string {
	return field.Value
}
