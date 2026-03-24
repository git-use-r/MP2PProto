package connection

import (
	entity "MP2PProto/transport/Entity"
	"net"
)

type Listener struct {
	Connect *net.UDPConn
	LAddr   *net.UDPAddr
}

func NewListener(localAddr string) (Listener, error) {

	lAddr, err1 := net.ResolveUDPAddr("udp", localAddr)
	if err1 != nil {
		return Listener{}, err1
	}

	conn, err3 := net.ListenUDP("udp", lAddr)

	return Listener{
		conn,
		lAddr,
	}, err3

}

func CloseListener(listener *Listener) error {
	err := listener.Connect.Close()
	if err != nil {
		return err
	}
	return nil
}

func ReadToListener(listener *Listener) (entity.Package, error) {
	return entity.Package{}, nil
}
