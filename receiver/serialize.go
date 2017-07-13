package receiver

import (
	"bytes"
	"encoding/binary"
)

func encodeInt64(t int64, buf []byte) error {
	b := bytes.NewBuffer(buf)
	b.Reset()
	err := binary.Write(b, binary.LittleEndian, t)
	return err
}
