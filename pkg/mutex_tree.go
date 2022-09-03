package pkg

import (
	"sync"
)

type MuTree struct {
	t *Tree
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
