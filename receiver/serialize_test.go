package receiver

import (
	"bytes"
	"encoding/binary"
	"testing"
	"time"

	"github.mit.edu/hari/nimbus-cc/packetops"
)

type Packet struct {
	Echo     int64  // time at which packet was sent
	RecvTime int64  // time packet reached receiver
	Payload  string // payload (useless)
}

func (pkt *Packet) Encode(
	size int,
) (*packetops.RawPacket, error) {
	return nil, nil
}

func (pkt *Packet) Decode(
	r *packetops.RawPacket,
) error {
	buf := r.Buf[8:]
	b := bytes.NewBuffer(buf)
	binary.Read(b, binary.LittleEndian, &pkt.RecvTime)

	return nil
}

func TestEncodeRecvTime(t *testing.T) {
	headerOffset = 8
	hdr := bytes.Repeat([]byte{0}, 16)
	ack := packetops.RawPacket{Buf: hdr}
	n := time.Now().UnixNano()
	makeAck(&ack, n)

	dec := Packet{}
	err := dec.Decode(&ack)
	if err != nil {
		t.Error(err)
	}

	if dec.RecvTime != n {
		t.Error("encoded incorrectly", dec.RecvTime, n, dec, ack.Buf)
	}
}

// benchmark how much time it takes to modify the packet
func BenchmarkRecvTime(b *testing.B) {
	hdr := bytes.Repeat([]byte{0}, 16)
	ack := packetops.RawPacket{Buf: hdr}
	for i := 0; i < b.N; i++ {
		makeAck(&ack, time.Now().UnixNano())
	}
}
