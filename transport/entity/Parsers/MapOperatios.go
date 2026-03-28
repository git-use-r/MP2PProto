package parsers

import (
	"github.com/malchik-one/MP2PProto/transport/entity/operations"
	"github.com/malchik-one/MP2PProto/transport/entity/operations/ByteType"
)

var bytSen = &ByteType.ByteSend{}

var MapOperationsString = map[string]operations.Operation{
	"ByteSend": bytSen,
}

var MapOperationsByte = map[byte]operations.Operation{
	5: bytSen,
}
