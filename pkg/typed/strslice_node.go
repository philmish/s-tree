package typed

import (
	"bytes"
	"encoding/gob"
)

func encodeStrSlice(value []string) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeStrSlice(value []byte) ([]string, error) {
	var res []string
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}

func StrSliceNode(value []string, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeStrSlice(value)
	if err != nil {
		return nil, err
	}
	n := TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     STRSLICE,
		Children: make([]*TypedNode, 0),
	}
	if parent != nil {
		parent.Children = append(parent.Children, &n)
	}
	return &n, err
}
