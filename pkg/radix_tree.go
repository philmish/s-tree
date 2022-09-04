package pkg

import (
	"strings"
	"sync"
)

type RadixTree struct {
	t *Tree
}

func NewRadix() *RadixTree {
	return &RadixTree{
		t: createTree(),
	}
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
	return rt.t.radixAdd(vals)
}

func (rt *RadixTree) AsyncPush(data string, wg *sync.WaitGroup) error {
	rt.t.Lock()
	defer func() {
		rt.t.Unlock()
		wg.Done()
	}()
	vals := valsFromString(data)
	err := rt.t.radixAdd(vals)
	return err
}

func (rt *RadixTree) Search(data string) error {
	vals := valsFromString(data)
	return rt.t.SearchSequence(vals)
}

func (rt *RadixTree) AsyncSearch(data string, wg *sync.WaitGroup) error {
	rt.t.Lock()
	defer func() {
		rt.t.Unlock()
		wg.Done()
	}()
	vals := valsFromString(data)
	err := rt.t.SearchSequence(vals)
	return err
}
