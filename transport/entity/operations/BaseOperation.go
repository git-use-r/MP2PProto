package operations

import (
	"net/netip"
)

type BaseOperation struct {
	Operation
	IndexPack uint32
	Sender    netip.AddrPort
	Recipient netip.AddrPort
}

































