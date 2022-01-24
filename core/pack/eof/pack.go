package eof

type EofP struct{}

func NewPack() *EofP {
	return &EofP{}
}

func (p *EofP) Pack(data interface{}) ([]byte, error) {
	return data.([]byte), nil
}

func (p *EofP) Unpack(data []byte) (interface{}, error) {
	return data, nil
}
