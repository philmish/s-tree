package kvdb

import (
	"fmt"
	"sync"

	"github.com/philmish/s-tree/pkg"
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
	kv.RLock()
	defer func() {
		kv.RUnlock()
		wg.Done()
	}()
	kBytes := []byte(key)
	node := kv.Levels[0].GetNodeByValue(kBytes)
	if node == nil {
		return "", fmt.Errorf("%s not set\n", key)
	}
	if len(node.Children) == 0 {
		return "", fmt.Errorf("%s not set\n", key)
	}
	res := string(node.Children[0].Value)
	return res, nil
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
