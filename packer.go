package giso8583

type Packer interface {
	Pack() ([]byte, error)
	Unpack() (Packer, error)
}
