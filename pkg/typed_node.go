package pkg

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Supported Types
var (
	STR      byte = byte(0)
	INT           = byte(1)
	BOOL          = byte(2)
	STRSLICE      = byte(3)
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
	default:
		return -1, fmt.Errorf("Unknown type code %d", t)
	}
}

func StrNode(value string, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeStr(value)
	return &TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     byte(0),
		Children: make([]*TypedNode, 0),
	}, err
}

func IntNode(value int, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeInt(value)
	return &TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     byte(1),
		Children: make([]*TypedNode, 0),
	}, err
}

func BoolNode(value bool, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeBool(value)
	return &TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     byte(2),
		Children: make([]*TypedNode, 0),
	}, err
}

func encodeStr(value string) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeStr(value []byte) (string, error) {
	var res string
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}

func encodeInt(value int) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeInt(value []byte) (int, error) {
	var res int
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}

func encodeBool(value bool) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeBool(value []byte) (bool, error) {
	var res bool
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}
