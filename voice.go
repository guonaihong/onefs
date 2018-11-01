package onefs

import (
	"bytes"
	"encoding/binary"
	"github.com/guonaihong/onefs/voice"
	"unsafe"
)

var ECMATable = crc64.MakeTable(crc64.ECMA)

const voiceWord = 8

type voice struct {
	voice.Voice
	paddingLen uint8
	crc        uint64
}

func (v *voice) marshal() (data []byte, err error) {
	hasher := crc64.New(ECMATable)
	padding := (v.XXX_Size() + unsafe.Sizeof(v.crc)) % voiceWord

	buf := &bytes.Buffer

	data2, err2 := proto.Marshal(&v.Voice)
	if err != nil {
		return nil, err2
	}

	data2 = append(data2, make([]byte, padding))

	buf.Write(data2)
	hasher.Write(data2)

	err = binary.Write(buf, binary.LittleEndian, hasher.Sum64())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (v *voice) unmarshal() {
}
