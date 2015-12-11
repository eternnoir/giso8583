package anfield

import (
	"unicode/utf8"
)

type ANField struct {
	Length int
	Value  string
}

func New() *ANField {
	return &ANField{}
}

func (field *ANField) Pack() ([]byte, error) {
	result := []byte(field.Value)
	field.Length = utf8.RuneCountInString(field.Value)
	return result, nil
}

func (field *ANField) Unpacke(byteary []byte) (*ANField, error) {
	field.Value = string(byteary)
	field.Length = utf8.RuneCountInString(field.Value)
	return field, nil
}
