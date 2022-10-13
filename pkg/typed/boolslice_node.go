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
	if err != nil {
		return nil, err
	}
	n := TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     BOOLSLICE,
		Children: make([]*TypedNode, 0),
	}
	if parent != nil {
		parent.Children = append(parent.Children, &n)
	}
	return &n, err
}
