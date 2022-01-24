package common

type CmP struct{}

func NewPack() *CmP {
	return &CmP{}
}

func (p *CmP) Pack(data interface{}) ([]byte, error) {
	return data.([]byte), nil
}

func (p *CmP) Unpack(data []byte) (interface{}, error) {
	return data, nil
}
