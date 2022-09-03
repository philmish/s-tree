package pkg


import (
	"testing"
)

func TestAddBranch(t *testing.T) {
	tree := createTree()
    vals := values{[]byte("abc"), []byte("de"),} 
	err := tree.addBranch(vals)
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	if tree.Depth() != 2 {
		t.Errorf("Expected 2 levels, found %d\n", tree.Depth())
	}
}

func TestSearchNode(t *testing.T) {
    tree := createTree()
    vals := values{[]byte("abc"), []byte("de"),} 
	err := tree.addBranch(vals)
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
    _, err = tree.SearchNode([]byte("de"))
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }
}
