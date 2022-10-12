package typed

import (
	"bytes"
	"encoding/gob"
)

func encodeBoolSlice(value []bool) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeBoolSlice(value []byte) ([]bool, error) {
	var res []bool
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}

func BoolSliceNode(value []bool, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeBoolSlice(value)
	return &TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     byte(5),
		Children: make([]*TypedNode, 0),
	}, err
}
