package typed

import (
	"fmt"
	"sync"
)

type TypedTree struct {
	sync.RWMutex
	Root   *TypedNode
	Levels []*TypedTreeLevel
}

func (tt *TypedTree) Leafs() []*TypedNode {
	var res []*TypedNode
	for _, lvl := range tt.Levels {
		res = append(res, lvl.getLeafs()...)
	}
	return res
}

func (tt *TypedTree) Depth() int {
	return len(tt.Levels)
}

func newTree() *TypedTree {
	return &TypedTree{
		Root:   nil,
		Levels: make([]*TypedTreeLevel, 0),
	}
}

func (tt *TypedTree) AddNodeByCords(value interface{}, level, branch int) error {
	if len(tt.Levels) < level || len(tt.Levels[level].nodes) < branch {
		return fmt.Errorf("Target node does not exist")
	}
	parent := tt.Levels[level].getBranch(branch)
	_, err := NewTypedNode(value, parent)
	return err
}

func (tt *TypedTree) AddBranch(values []interface{}) error {
	tt.RLock()
	defer tt.RUnlock()
	n, err := NewTypedNode(values[0], tt.Root)
	if err != nil {
		return err
	}
	if tt.Root == nil {
		tt.Root = n
		tt.Levels = append(tt.Levels, newLevel())
	}
	currLvl := 0
	for _, val := range values[1:] {
		n, err = NewTypedNode(val, n)
		if err != nil {
			return err
		}
		tt.Levels[currLvl].nodes = append(tt.Levels[currLvl].nodes, n)
		if currLvl+1 == len(tt.Levels) {
			tt.Levels = append(tt.Levels, newLevel())
		}
		currLvl++
	}
	return nil
}
