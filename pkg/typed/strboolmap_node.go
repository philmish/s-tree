package typed

import (
	"bytes"
	"encoding/gob"
)

func encodeStrBoolMap(value map[string]bool) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	return buf.Bytes(), err
}

func decodeStrBoolMap(value []byte) (map[string]bool, error) {
	var res map[string]bool
	bf := bytes.NewBuffer(value)
	err := gob.NewDecoder(bf).Decode(&res)
	return res, err
}

func StrBoolMapNode(value map[string]bool, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeStrBoolMap(value)
	if err != nil {
		return nil, err
	}
	n := TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     STRBOOLMAP,
		Children: make([]*TypedNode, 0),
	}
	if parent != nil {
		parent.Children = append(parent.Children, &n)
	}
	return &n, err
}
