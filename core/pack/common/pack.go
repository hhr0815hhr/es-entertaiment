package common

type CmP struct{}

func NewPack() *CmP {
	return &CmP{}
}

/**
 * 固定包头包长+包体
 * 包头：4字节 + 1字节
 */

func (p *CmP) Pack(data interface{}) ([]byte, error) {
	// length := len(data.([]byte))
	// buf := make([]byte, length+4)

	return data.([]byte), nil
}

func (p *CmP) Unpack(data []byte) (interface{}, error) {

	// head := string(data[:4])
	// size := int(data[4])
	// body := data[5:]
	return data, nil
}
