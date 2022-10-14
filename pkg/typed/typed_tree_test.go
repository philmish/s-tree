package typed

import (
	"testing"
)

func TestAddBranch(t *testing.T) {
	tree := newTree()
	vals := []interface{}{
		"Hello",
		1,
		true,
		[]string{"Wo", "rld"},
	}
	err := tree.AddBranch(vals)
	if err != nil {
		t.Errorf("%s\nFailed to add branch", err.Error())
	}
	if tree.Depth() != 4 {
		t.Errorf("Expected depth of 4, found %d", tree.Depth())
	}
}

func TestLeafs(t *testing.T) {
	tree := newTree()
	vals1 := []interface{}{
		"Hello",
		1,
		true,
		[]string{"Wo", "rld"},
	}
	vals2 := []interface{}{
		"A",
		true,
	}
	err := tree.AddBranch(vals1)
	if err != nil {
		t.Errorf("%s\nFailed to add first branch", err.Error())
	}
	err = tree.AddBranch(vals2)
	if err != nil {
		t.Errorf("%s\nFailed to add second branch", err.Error())
	}
	leafs := tree.Leafs()
	for _, l := range leafs {
		v, err := l.GetValue()
		if err != nil {
			t.Errorf("Failed to read node value of leaf: %s", err.Error())
		}
		t.Logf("Node value: %v", v)
	}
	if len(leafs) != 2 {
		t.Errorf("Expected 2 leafs found %d", len(leafs))
	}
}

func TestMerge(t *testing.T) {
	tree1 := newTree()
	tree2 := newTree()
	vals1 := []interface{}{
		"Hello",
		1,
		true,
		[]string{"Wo", "rld"},
	}
	vals2 := []interface{}{
		"A",
		true,
	}
	err := tree1.AddBranch(vals1)
	if err != nil {
		t.Errorf("%s\nFailed to add branch to first tree", err.Error())
	}
	err = tree2.AddBranch(vals2)
	if err != nil {
		t.Errorf("%s\nFailed to add branch to second tree", err.Error())
	}
	err = tree1.Merge(tree2, 2, 0)
	if err != nil {
		t.Errorf("%s\nFailed to Merge trees", err.Error())
	}
	if tree1.Depth() != 6 {
		t.Errorf("Expected depth of 6 found %d", tree1.Depth())
	}

}
