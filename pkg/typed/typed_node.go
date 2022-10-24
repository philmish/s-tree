package typed

import (
	"bytes"
	"fmt"
)

// Max sizes for map and slice type data
var (
	MAXSLICE = 1024
	MAXMAP   = 1024
)

// Supported Types
var (
	STR        = byte(0)
	INT        = byte(1)
	BOOL       = byte(2)
	STRSLICE   = byte(3)
	INTSLICE   = byte(4)
	BOOLSLICE  = byte(5)
	STRSTRMAP  = byte(6)
	STRINTMAP  = byte(7)
	STRBOOLMAP = byte(8)
)

type TypedNode struct {
	Parent   *TypedNode
	Value    []byte
	Type     byte
	Children []*TypedNode
}

func (n *TypedNode) IsEqualTo(node *TypedNode) bool {
	return n.Type == node.Type && bytes.Compare(n.Value, node.Value) == 0
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
	case STRSTRMAP:
		return decodeStrStrMap(n.Value)
	case STRINTMAP:
		return decodeStrIntMap(n.Value)
	case STRBOOLMAP:
		return decodeStrBoolMap(n.Value)
	default:
		return -1, fmt.Errorf("Unknown type code %d", t)
	}
}

func NewTypedNode(value interface{}, parent *TypedNode) (*TypedNode, error) {
	switch value.(type) {
	case string:
		return StrNode(value.(string), parent)
	case int:
		return IntNode(value.(int), parent)
	case bool:
		return BoolNode(value.(bool), parent)
	case []string:
		return StrSliceNode(value.([]string), parent)
	case []int:
		return IntSliceNode(value.([]int), parent)
	case []bool:
		return BoolSliceNode(value.([]bool), parent)
	case map[string]string:
		return StrStrMapNode(value.(map[string]string), parent)
	case map[string]int:
		return StrIntMapNode(value.(map[string]int), parent)
	case map[string]bool:
		return StrBoolMapNode(value.(map[string]bool), parent)
	default:
		return nil, fmt.Errorf("Unsupported data type")
	}
}
