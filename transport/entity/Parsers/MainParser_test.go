package parsers

import (
	"testing"

	"github.com/malchik-one/MP2PProto/transport/entity/operations"
	"github.com/malchik-one/MP2PProto/transport/entity/operations/ByteType"
)

var b1 = ByteType.ByteSend{
	TagSend: 4,
}

var b2 = ByteType.ByteSend{
	TagSend: 2,
}

var b3 = ByteType.ByteSend{
	TagSend: 7,
}

func TestMainParseToBinary(t *testing.T) {
	b, err := MainParseToBinary([]operations.Operation{
		&b1,
		&b2,
		&b3,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range b {
		println(i)
	}
}

func TestMainParseToStruct(t *testing.T) {
	b := []byte{5, 5, 5, 5, 5}

	ops, err := MainParseToStruct(b)
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range ops {
		println(i.GetTag())
	}
}
