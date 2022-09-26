package pkg

import (
	"fmt"
	"sync"
)

type value []byte
type values [][]byte

type Tree struct {
	sync.RWMutex
	Root   *Node
	Levels []*TreeLevel
}

func (t *Tree) Depth() int {
	return len(t.Levels)
}

func (t *Tree) Leafs() []*Node {
	res := make([]*Node, 0)
	for _, lvl := range t.Levels {
		res = append(res, lvl.getLeafs()...)
	}
	return res
}

func createTree() (tree *Tree) {
	root := newNode([]byte(""))
	return &Tree{
		Root:   root,
		Levels: make([]*TreeLevel, 0),
	}
}

func NewTree(rootValue string) *Tree {
	r := []byte(rootValue)
	return &Tree{
		Root:   newNode(r),
		Levels: make([]*TreeLevel, 0),
	}
}

func (t *Tree) lastLevel() (*TreeLevel, error) {
	if t.Depth() == 0 {
		return nil, fmt.Errorf("Empty Tree")
	}
	return t.Levels[t.Depth()-1], nil
}

func (t *Tree) addBranch(data values) error {
	cursor := t.Root
	var b []byte
	var n *Node
	var err error
	currLvl := 1
	for len(data) > 0 {
		b, data = data[0], data[1:]
		n = CreateNode(b, cursor)
		if t.Depth() < currLvl {
			err = t.addLevel(n)
			if err != nil {
				return err
			}
		} else {
			err = t.Levels[currLvl-1].Append(n)
			if err != nil {
				return err
			}
		}
		currLvl++
		cursor = n
	}
	return nil
}

func (t *Tree) addLevel(n *Node) error {
	lvl := newLevel()
	if t.Depth() == 0 {
		lvl.Nodes = append(lvl.Nodes, n)
		n.setParent(t.Root)
		t.Root.addChild(n)
		t.Levels = append(t.Levels, lvl)
		return nil
	}
	if n.Parent == nil {
		return fmt.Errorf("Node with value: %s has no parent", n.Value)
	}
	lvl.Nodes = append(lvl.Nodes, n)
	t.Levels = append(t.Levels, lvl)
	return nil
}

func (t *Tree) SearchNode(val []byte) (*Node, error) {
	if t.Depth() == 0 {
		return nil, fmt.Errorf("Tree is emtpy")
	}
	for _, lvl := range t.Levels {
		exists := lvl.GetNodeByValue(val)
		if exists != nil {
			return exists, nil
		}
	}
	return nil, fmt.Errorf("%s not found\n", string(val))
}

func (t *Tree) SearchSequence(vals values) error {
	if len(vals) > t.Depth() {
		return fmt.Errorf("%v not found\n", vals)
	} else if t.Depth() == 0 {
		return fmt.Errorf("Tree is empty\n")
	}
	var b value
	cursor := t.Root
	for len(vals) > 0 {
		b, vals = vals[0], vals[1:]
		found := false
		for _, c := range cursor.Children {
			if c.CompareVal(b) {
				cursor = c
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("%v not found\n", string(b))
		}
	}
	return nil
}

func (t *Tree) GetLevelValues(lvl int) ([]string, error) {
	if t.Depth() == 0 {
		return []string{}, fmt.Errorf("Tree is empty")
	}
	if lvl > t.Depth() {
		return []string{}, fmt.Errorf("Level not in tree")
	}
	values := make([]string, 0)
	for _, n := range t.Levels[lvl-1].Nodes {
		values = append(values, string(n.Value))
	}
	return values, nil
}

func (t *Tree) ThreadSafeAddBranch(vals values) error {
	t.Lock()
	defer func() {
		t.Unlock()
	}()
	err := t.addBranch(vals)
	return err
}

func (t *Tree) ThreadSafeRadixAdd(vals values) error {
	t.Lock()
	defer func() {
		t.Unlock()
	}()
	err := t.radixAdd(vals)
	return err
}

func (t *Tree) ThreadSafeSearchNode(val []byte) (*Node, error) {
	t.RLock()
	defer func() {
		t.RUnlock()
	}()
	res, err := t.SearchNode(val)
	return res, err
}

func (t *Tree) ThreadSafeSearchSeq(vals values) error {
	t.RLock()
	defer func() {
		t.RUnlock()
	}()
	err := t.SearchSequence(vals)
	return err
}

func (t *Tree) ThreadSafeGetLevelValues(lvl int) ([]string, error) {
	t.RLock()
	defer func() {
		t.RUnlock()
	}()
	vals, err := t.GetLevelValues(lvl)
	return vals, err
}
