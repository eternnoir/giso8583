package nfield

import (
	"encoding/hex"
	"fmt"
	"unicode/utf8"

	"github.com/eternnoir/giso8583"
)

type NField struct {
	Length int
	Value  string
}

func New() *NField {
	return new(NField)
}

func (field *NField) Pack() ([]byte, error) {
	ret, err := hex.DecodeString(field.Value)
	if err != nil {
		return nil, fmt.Errorf("NField Pack error. Value:%s. %s", field.Value, err)
	}
	field.Length = utf8.RuneCountInString(field.Value)
	return ret, nil
}

func (field *NField) Unpack(byteary []byte) (giso8583.Packer, error) {
	field.Value = hex.EncodeToString(byteary)
	field.Length = utf8.RuneCountInString(field.Value)
	return field, nil
}

func (field *NField) GetValue() string {
	return field.Value
}
