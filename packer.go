package iso8583

type Packer interface {
	Pack() []byte
	Unpack() Packer
}
