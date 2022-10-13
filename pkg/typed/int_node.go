package typed

import (
	"bytes"
	"encoding/gob"
)

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

func IntNode(value int, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeInt(value)
	if err != nil {
		return nil, err
	}
	n := TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     INT,
		Children: make([]*TypedNode, 0),
	}
	if parent != nil {
		parent.Children = append(parent.Children, &n)
	}
	return &n, err
}
