package connection

import (
	"net"
)

type Connector struct {
	Connect *net.UDPConn
	LAddr   *net.UDPAddr
	RAddr   *net.UDPAddr
}

func NewConnect(localAddr string, remoteAddr string) (Connector, error) {

	lAddr, err1 := net.ResolveUDPAddr("udp", localAddr)
	if err1 != nil {
		return Connector{}, err1
	}

	rAddr, err2 := net.ResolveUDPAddr("udp", remoteAddr)
	if err2 != nil {
		return Connector{}, err2
	}

	conn, err3 := net.DialUDP("udp", lAddr, rAddr)

	return Connector{
		conn,
		lAddr,
		rAddr,
	}, err3

}

func CloseConnect(connector *Connector) error {
	err := connector.Connect.Close()
	if err != nil {
		return err
	}
	return nil
}

func WriteToConnect(connector *Connector, message string) error {
	return nil
}


























