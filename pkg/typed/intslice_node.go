package typed

import (
	"bytes"
	"encoding/gob"
)

func encodeIntSlice(value []int) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeIntSlice(value []byte) ([]int, error) {
	var res []int
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}

func IntSliceNode(value []int, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeIntSlice(value)
	return &TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     byte(4),
		Children: make([]*TypedNode, 0),
	}, err
}
