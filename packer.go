package giso8583

type Packer interface {
	Pack() ([]byte, error)
	Unpack([]byte) (Packer, error)
	GetValue() string
}
