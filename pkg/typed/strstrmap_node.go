package typed

import (
	"bytes"
	"encoding/gob"
)

func encodeStrStrMap(value map[string]string) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeStrStrMap(value []byte) (map[string]string, error) {
	var res map[string]string
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}

func StrStrMapNode(value map[string]string, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeStrStrMap(value)
	if err != nil {
		return nil, err
	}
	n := TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     STRSTRMAP,
		Children: make([]*TypedNode, 0),
	}
	if parent != nil {
		parent.Children = append(parent.Children, &n)
	}
	return &n, err
}
