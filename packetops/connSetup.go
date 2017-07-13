package packetops

import (
	"fmt"
	"net"
)

func SetupClientSock(ip string, port string) (*net.UDPConn, *net.UDPAddr, error) {
	addr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return nil, nil, err
	}
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return nil, nil, err
	}

	return conn, addr, nil
}

func SetupListeningSock(port string) (*net.UDPConn, *net.UDPAddr, error) {
	// set up syn listening socket
	addr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%s", port))
	if err != nil {
		return nil, nil, err
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return nil, nil, err
	}

	return conn, addr, nil
}

func SendSyn(conn *net.UDPConn, syn Packet) error {
	err := SendAck(conn, syn)
	if err != nil {
		return err
	}

	return nil
}

// wait for the SYN
func ListenForSyn(
	conn *net.UDPConn,
	listenAddr *net.UDPAddr,
	p Packet,
) (*net.UDPConn, error) {
	fromAddr, err := RecvPacket(conn, p)
	if err != nil {
		return nil, err
	}

	// close and reopen
	conn.Close()

	// dial connection to send ACKs
	newConn, err := net.DialUDP("udp4", listenAddr, fromAddr)
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to ", fromAddr)
	return newConn, nil
}

// send the given syn
// receive the synack into the same packet buffer
func SynAckExchange(conn *net.UDPConn, expSrc *net.UDPAddr, syn Packet) error {
	err := SendSyn(conn, syn)
	if err != nil {
		return err
	}

	srcAddr, err := RecvPacket(conn, syn)
	if err != nil {
		return err
	}
	if srcAddr.String() != expSrc.String() {
		return fmt.Errorf("got packet from unexpected src: %s; expected %s", srcAddr, expSrc)
	}

	return nil
}
