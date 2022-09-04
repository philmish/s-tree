package pkg

import (
	"strings"
	"sync"
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

func (rt *RadixTree) Push(data string) error {
	vals := valsFromString(data)
	return rt.radixAdd(vals)
}

func (rt *RadixTree) Search(data string) error {
	vals := valsFromString(data)
	return rt.SearchSequence(vals)
}

func (rt *RadixTree) TSafePush(data string, wg *sync.WaitGroup) error {
	vals := valsFromString(data)
	err := rt.ThreadSafeRadixAdd(vals, wg)
	return err
}

func (rt *RadixTree) TSafeSearch(data string, wg *sync.WaitGroup) error {
	vals := valsFromString(data)
	err := rt.ThreadSafeSearchSeq(vals, wg)
	return err
}
