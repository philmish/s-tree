package kvdb

import (
	"sync"
	"testing"
)

func TestStoreSet(t *testing.T) {
	store := NewStore()
	data := [][]string{
		{"a", "b"},
		{"a", "c"},
	}

	wg := sync.WaitGroup{}
	for _, val := range data {
		wg.Add(1)
		go func(v []string, wg *sync.WaitGroup) {
			store.Set(v[0], v[1], wg)
		}(val, &wg)
	}
	wg.Wait()
	if store.Depth() != 2 {
		t.Errorf("Expected depth 2 found %d\n", store.Depth())
	}
	lenKeys := store.Levels[0].Length()
	if lenKeys != 1 {
		t.Errorf("Expected 1 key found %d\n", lenKeys)
	}
}
