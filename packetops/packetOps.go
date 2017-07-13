package packetops

import (
	"net"
)

// helper function to encode and send a packet
func SendPacket(
	conn *net.UDPConn,
	pkt Packet,
	size int,
) error {
	rp, err := pkt.Encode(size)

	_, err = conn.Write(rp.Buf)
	if err != nil {
		return err
	}

	return nil
}

func SendAck(
	conn *net.UDPConn,
	pkt Packet,
) error {
	return SendPacket(conn, pkt, 0)
}

// helper function to receive and decode a packet
// read packet written to p
func RecvPacket(
	conn *net.UDPConn,
	p Packet,
) (*net.UDPAddr, error) {
	rcvd := &RawPacket{Buf: make([]byte, 1500)}
	err := Listen(conn, rcvd)
	if err != nil {
		return nil, err
	}

	err = p.Decode(rcvd)
	if err != nil {
		return nil, err
	}

	return rcvd.From, err
}
