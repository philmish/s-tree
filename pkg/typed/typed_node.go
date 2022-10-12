package typed

import (
	"fmt"
)

// Sizes
var (
	MAXSLICE = 1024
	MAXMAP   = 1024
)

// Supported Types
var (
	STR       = byte(0)
	INT       = byte(1)
	BOOL      = byte(2)
	STRSLICE  = byte(3)
	INTSLICE  = byte(4)
	BOOLSLICE = byte(5)
	// To be implemented
	STRSTRMAP = byte(6)
	STRINTMAP = byte(7)
)

type TypedNode struct {
	Parent   *TypedNode
	Value    []byte
	Type     byte
	Children []*TypedNode
}

func (n *TypedNode) GetValue() (interface{}, error) {
	t := n.Type
	switch t {
	case STR:
		return decodeStr(n.Value)
	case INT:
		return decodeInt(n.Value)
	case BOOL:
		return decodeBool(n.Value)
	case STRSLICE:
		return decodeStrSlice(n.Value)
	case INTSLICE:
		return decodeIntSlice(n.Value)
	case BOOLSLICE:
		return decodeBoolSlice(n.Value)
	default:
		return -1, fmt.Errorf("Unknown type code %d", t)
	}
}
