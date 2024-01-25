package datafield

import (
	"github.com/bhagyaraj1208117/andes-abc-1/marshal"
)

// ArgsOperationDataFieldParser holds all the components required to create a new instance of data field parser
type ArgsOperationDataFieldParser struct {
	AddressLength int
	Marshalizer   marshal.Marshalizer
}
