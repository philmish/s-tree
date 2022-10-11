package typed

import (
	"bytes"
	"encoding/gob"
)

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

func StrNode(value string, parent *TypedNode) (*TypedNode, error) {
	val, err := encodeStr(value)
	return &TypedNode{
		Parent:   parent,
		Value:    val,
		Type:     byte(0),
		Children: make([]*TypedNode, 0),
	}, err
}
