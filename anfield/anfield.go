package anfield

import (
	"github.com/eternnoir/giso8583"
	"unicode/utf8"
)

type ANField struct {
	Length int
	Value  string
}

func New() *ANField {
	return new(ANField)
}

func (field *ANField) Pack() ([]byte, error) {
	result := []byte(field.Value)
	field.Length = utf8.RuneCountInString(field.Value)
	return result, nil
}

func (field *ANField) Unpack(byteary []byte) (giso8583.Packer, error) {
	field.Value = string(byteary)
	field.Length = utf8.RuneCountInString(field.Value)
	return field, nil
}
