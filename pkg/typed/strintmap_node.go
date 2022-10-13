package typed

import (
	"bytes"
	"encoding/gob"
)

func encodeStrIntMap(value map[string]int) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeStrIntMap(value []byte) (map[string]int, error) {
	var res map[string]int
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}

func StrIntMapNode(value map[string]int, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeStrIntMap(value)
	if err != nil {
		return nil, err
	}
	n := TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     STRINTMAP,
		Children: make([]*TypedNode, 0),
	}
	if parent != nil {
		parent.Children = append(parent.Children, &n)
	}
	return &n, err
}
