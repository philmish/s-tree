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

func TestStoreGet(t *testing.T) {
	store := NewStore()
	data := [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	wg := sync.WaitGroup{}
	for _, val := range data {
		wg.Add(1)
		go func(v []string, wg *sync.WaitGroup) {
			store.Set(v[0], v[1], wg)
		}(val, &wg)
	}
	wg.Wait()
	wg.Add(2)
	v, err := store.Get("a", &wg)
	v2, err2 := store.Get("b", &wg)
	wg.Wait()
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	if err2 != nil {
		t.Errorf("%s\n", err2.Error())
	}
	if v != "b" {
		t.Errorf("Expected a to have value b got %s\n", v)
	}
	if v2 != "c" {
		t.Errorf("Expected b to have value c got %s\n", v2)
	}
}
