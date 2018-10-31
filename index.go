package onefs

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
)

const (
	notUse       = math.MaxUint32
	indexSize    = 16
	indexVersion = "v0.0.1"
)

type indexVersion struct {
	Version [indexSize * 2]byte
}

func indexVersionNew() *indexVersion {
	v := &indexVersion{}
	copy(v.version, indexVersion)
	return v
}

func (iv *indexVersion) marshal() (data []byte, err error) {
	buf := &bytes.Buffer{}
	err = binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		return
	}

	data = buf.Bytes()
}

func (iv *indexVersion) unmarshal(data []byte) (err error) {
	err = binary.Read(bytes.NewBuffer(data), binary.LittleEndian, iv)
	if err != nil {
		return err
	}

	//todo: 比较版本号，高版本兼容低版本，反之不行
	if indexVersion != string(iv.Version) {
		panic("unkown version:" + iv.Version)
	}

	return nil
}

type index struct {
	Key    uint64
	Offset uint32
	Size   uint32
}

func (i *index) marshal() (data []byte, err error) {
	buf := &bytes.Buffer{}
	err = binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		return
	}

	data = buf.Bytes()
	return
}

func (i *index) unmarshal(data []byte) error {
	if len(data) != indexSize {
		panic("Wrong index length")
	}

	return binary.Read(bytes.NewBuffer(data), binary.LittleEndian, i)
}
