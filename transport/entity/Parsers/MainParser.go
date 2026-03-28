package parsers

import (
	"github.com/malchik-one/MP2PProto/transport/entity/operations"
)

func MainParseToBinary(sliceOp []operations.Operation) ([]byte, error) {

	var b []byte = make([]byte, len(sliceOp))

	for l, i := range sliceOp {
		marshBinar, err := i.MarshalBinary()
		if err != nil {
			return []byte{}, err
		}
		copy(b[l:], marshBinar)
	}

	return b, nil
}

func MainParseToStruct(b []byte) ([]operations.Operation, error) {

	var ops []operations.Operation = make([]operations.Operation, len(b))

	for l, i := range b {
		ops[l] = MapOperationsByte[i]
		_, err := ops[l].UnmarshalBinary(b)
		if err != nil {
			return nil, err
		}
	}

	return ops, nil

}
