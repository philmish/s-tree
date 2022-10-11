package typed

import (
	"bytes"
	"encoding/gob"
)

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

func BoolNode(value bool, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeBool(value)
	return &TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     byte(2),
		Children: make([]*TypedNode, 0),
	}, err
}
