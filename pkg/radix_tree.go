package pkg

import (
	"strings"
)

type RadixTree struct {
	*Tree
}

func NewRadix() *RadixTree {
	t := RadixTree{}
	t.Tree = createTree()
	return &t
}

func valsFromString(data string) values {
	symbols := strings.Split(data, "")
	vals := make(values, 0)
	var b []byte
	for len(symbols) > 0 {
		b, symbols = []byte(symbols[0]), symbols[1:]
		vals = append(vals, b)
	}
	return vals
}

func (t *Tree) radixAdd(vals values) error {
	if t.Depth() == 0 {
		return t.addBranch(vals)
	}
	cursor := t.Root
	var lvlCursor *TreeLevel
	curLvl := 0
	var b value
	for len(vals) > 0 {
		lvlCursor = t.Levels[curLvl]
		b, vals = vals[0], vals[1:]
		exists := lvlCursor.GetNodeByValue(b)
		if exists == nil {
			n := CreateNode(b, cursor)
			lvlCursor.Nodes = append(lvlCursor.Nodes, n)
			cursor = n
		} else {
			cursor = exists
		}
		if len(vals) > 0 {
			curLvl += 1
		}
		if curLvl >= t.Depth() {
			t.Levels = append(t.Levels, newLevel())
		}
	}
	return nil
}

func (rt *RadixTree) PushStr(data string) error {
	vals := valsFromString(data)
	return rt.radixAdd(vals)
}

func (rt *RadixTree) SearchStr(data string) error {
	vals := valsFromString(data)
	return rt.SearchSequence(vals)
}

func (rt *RadixTree) TSafePushStr(data string) error {
	vals := valsFromString(data)
	err := rt.ThreadSafeRadixAdd(vals)
	return err
}

func (rt *RadixTree) TSafeSearchStr(data string) error {
	vals := valsFromString(data)
	err := rt.ThreadSafeSearchSeq(vals)
	return err
}
