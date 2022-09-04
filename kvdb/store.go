package kvdb

import (
	"github.com/philmish/s-tree/pkg"
	"sync"
)

type KvStore struct {
	*pkg.Tree
}

func NewStore() *KvStore {
	s := KvStore{}
	s.Tree = pkg.NewTree("")
	return &s
}

func (kv *KvStore) keyExists(key string) bool {
	kv.RLock()
	defer kv.RUnlock()
	if kv.Depth() == 0 {
		return false
	}
	node := kv.Levels[0].GetNodeByValue([]byte(key))
	return node != nil
}

func (kv *KvStore) Get(key string, wg *sync.WaitGroup) (string, error) {
	kBytes := []byte(key)
	keyNode, err := kv.ThreadSafeSearchNode(kBytes, wg)
	if err != nil {
		return "", err
	}
	return string(keyNode.Value), nil
}

func (kv *KvStore) Set(key, value string, wg *sync.WaitGroup) error {
	kBytes := []byte(key)
	vBytes := []byte(value)
	kExists := kv.keyExists(key)
	if kExists {
		node, err := kv.ThreadSafeSearchNode(kBytes, wg)
		if err != nil {
			return err
		}
		node.Children = make([]*pkg.Node, 0)
		_ = pkg.CreateNode(vBytes, node)
		return nil
	}
	vals := [][]byte{kBytes, vBytes}
	err := kv.ThreadSafeAddBranch(vals, wg)
	return err
}
