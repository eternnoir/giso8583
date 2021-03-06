package llvarzfield

import (
	"encoding/hex"
	"strconv"

	util "github.com/eternnoir/giso8583/util"
)

type LLVARZField struct {
	Length int
	Value  string
}

func New() *LLVARZField {
	return new(LLVARZField)
}

func (field *LLVARZField) Pack() ([]byte, error) {
	if (len(field.Value) % 2) == 1 {
		field.Value = field.Value + "0"
	}
	field.Length = len(field.Value)
	lengthHeader := strconv.Itoa(field.Length / 2)
	rawHexData := util.PadLeft(lengthHeader, "0", 2) + field.Value
	return hex.DecodeString(rawHexData)
}

func (field *LLVARZField) Unpack(byteary []byte) (*LLVARZField, error) {
	rawHex := hex.EncodeToString(byteary)
	withoutLenHex := rawHex[2:len(rawHex)]
	field.Value = withoutLenHex
	field.Length = len(field.Value)
	return field, nil
}

func (field *LLVARZField) GetValue() string {
	return field.Value
}
