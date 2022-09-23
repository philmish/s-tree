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

func TestGetLevelValues(t *testing.T) {
	tree := createTree()
	data := []values{
		{[]byte("a"), []byte("b")},
		{[]byte("c"), []byte("d")},
	}
	err := tree.addBranch(data[0])
	if err != nil {
		t.Errorf("Failed to add branch for level values test: %s\n", err.Error())
	}
	err = tree.addBranch(data[1])
	if err != nil {
		t.Errorf("Failed to add branch for level values test: %s\n", err.Error())
	}
	vals, err := tree.GetLevelValues(1)
	if err != nil {
		t.Errorf("Failed to fetch level values from lvl 1: %s\n", err.Error())
	}
	if len(vals) != 2 {
		t.Errorf("Expected 2 values in level 1, got %d", len(vals))
	}
	for _, v := range vals {
		if v != "a" && v != "c" {
			t.Errorf("Expectd a or c as value, got: %s", v)
		}
	}
}

func TestThreadSafeAdd(t *testing.T) {
	tree := createTree()
	data := []values{
		{[]byte("abc"), []byte("de")},
		{[]byte("abc"), []byte("ijk")},
	}
	wg := new(sync.WaitGroup)
	for _, i := range data {
		wg.Add(1)
		go func(n values) {
			err := tree.ThreadSafeAddBranch(n)
			if err != nil {
				fmt.Println("Err")
			}
			wg.Done()
		}(i)
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
	wg := new(sync.WaitGroup)
	for _, i := range data {
		wg.Add(1)
		go func(n values) {
			err := tree.ThreadSafeRadixAdd(n)
			if err != nil {
				fmt.Println("Err")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	if tree.Depth() != 4 {
		t.Errorf("Expected depth 4 got %d", tree.Depth())
	}
	if tree.Levels[1].Length() != 1 {
		t.Errorf("Expected level 2 to have length 1, found %d\n", tree.Levels[1].Length())
	}
}
