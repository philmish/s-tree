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

type treeCords struct {
	level  int
	branch int
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
