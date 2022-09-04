package pkg

import (
	"fmt"
	"sync"
	"testing"
)

func TestAddBranch(t *testing.T) {
	tree := createTree()
	vals := values{[]byte("abc"), []byte("de")}
	err := tree.addBranch(vals)
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	if tree.Depth() != 2 {
		t.Errorf("Expected 2 levels, found %d\n", tree.Depth())
	}
}

func TestMultipleAddBranch(t *testing.T) {
	tree := createTree()
	data := []values{
		{[]byte("abc"), []byte("de")},
		{[]byte("abc"), []byte("ijk")},
	}
	var err error
	for _, i := range data {
		err = tree.addBranch(i)
		if err != nil {
			t.Errorf("Failed to add values %v and %v\n", i[0], i[1])
		}
	}
	if tree.Depth() != 2 {
		t.Errorf("Expected depth 2, got %d\n", tree.Depth())
	}
}

func TestSearchNode(t *testing.T) {
	tree := createTree()
	vals := values{[]byte("abc"), []byte("de")}
	err := tree.addBranch(vals)
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	_, err = tree.SearchNode([]byte("de"))
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
}

func TestLeafs(t *testing.T) {
	tree := createTree()
	data := []values{
		{[]byte("abc"), []byte("de")},
		{[]byte("abc"), []byte("ijk")},
	}
	var err error
	for _, i := range data {
		err = tree.addBranch(i)
		if err != nil {
			t.Errorf("Failed to add values %v and %v\n", i[0], i[1])
		}
	}
	leafs := tree.Leafs()
	if len(leafs) != 2 {
		t.Errorf("Expected 2 leafs got %d", len(leafs))
	}
	for _, i := range leafs {
		if string(i.Value) != "de" && string(i.Value) != "ijk" {
			t.Errorf("Unexpected leaf: %s", string(i.Value))
		}
	}
}

func TestThreadSafeAdd(t *testing.T) {
	tree := createTree()
	data := []values{
		{[]byte("abc"), []byte("de")},
		{[]byte("abc"), []byte("ijk")},
	}
	wg := sync.WaitGroup{}
	for _, i := range data {
		wg.Add(1)
		go func(n values, wg *sync.WaitGroup) {
			err := tree.ThreadSafeAddBranch(n, wg)
			if err != nil {
				fmt.Println("Err")
			}
		}(i, &wg)
	}
	wg.Wait()
	if tree.Depth() != 2 {
		t.Logf("%v\n", tree.Levels[0].Length())
		t.Errorf("Expected depth 2 got %d", tree.Depth())
	}
}

func TestThreadSafeRadixAdd(t *testing.T) {
	tree := createTree()
	data := []values{
		{[]byte("a"), []byte("b"), []byte("c"), []byte("d")},
		{[]byte("a"), []byte("b"), []byte("i"), []byte("j")},
	}
	wg := sync.WaitGroup{}
	for _, i := range data {
		wg.Add(1)
		go func(n values, wg *sync.WaitGroup) {
			err := tree.ThreadSafeRadixAdd(n, wg)
			if err != nil {
				fmt.Println("Err")
			}
		}(i, &wg)
	}
	wg.Wait()
	if tree.Depth() != 4 {
		t.Errorf("Expected depth 4 got %d", tree.Depth())
	}
	if tree.Levels[1].Length() != 1 {
		t.Errorf("Expected level 2 to have length 1, found %d\n", tree.Levels[1].Length())
	}
}
