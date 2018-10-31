package onefs

type voice struct {
	voice  []byte
	pading []byte
	crc    uint64
}

func (v *voice) marshal() {
}

func (v *voice) unmarshal() {
}
