package ByteType

import (
	"errors"

	"github.com/malchik-one/MP2PProto/transport/entity/operations"
)

var _ operations.Operation = (*ByteSend)(nil)

type ByteSend struct {
	TagSend byte
}

func (b *ByteSend) MarshalBinary() ([]byte, error) {

	if b == nil {
		return []byte{}, errors.New("Cтруктура b не может быть равна nil")
	}

	return []byte{b.TagSend}, nil
}

func (b *ByteSend) UnmarshalBinary(data []byte) (n int, err error) {
	b.TagSend = 5
	return 1, nil
}

func (b *ByteSend) GetTag() (tagByt byte) {
	return b.TagSend
}
