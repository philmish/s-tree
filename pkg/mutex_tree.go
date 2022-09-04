package pkg

import (
	"sync"
)

type MuTree struct {
	t *Tree
}

func NewMuTree() *MuTree {
    return &MuTree{
        t: createTree(),
    }
}

func (mt *MuTree) Push(vals values, wg *sync.WaitGroup) error {
	mt.t.Lock()
	defer func() {
		mt.t.Unlock()
		wg.Done()
	}()
	err := mt.t.addBranch(vals)
	return err
}

func (mt *MuTree) Search(vals values, wg *sync.WaitGroup) error {
	mt.t.RLock()
	defer func() {
		mt.t.RUnlock()
		wg.Done()
	}()
	err := mt.t.SearchSequence(vals)
	return err
}

type MuRedixTree struct {
	t *RadixTree
}

func NewMuRTree() *MuRedixTree {
    return &MuRedixTree{
        t: NewRadix(),
    }
}

func (mrt *MuRedixTree) Push(vals string, wg *sync.WaitGroup) error {
	err := mrt.t.AsyncPush(vals, wg)
	return err
}

func (mrt *MuRedixTree) Search(vals string, wg *sync.WaitGroup) error {
	err := mrt.t.AsyncSearch(vals, wg)
	return err
}
