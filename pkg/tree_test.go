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

func TestMultipleAddBranch(t *testing.T) {
    tree :=  createTree()
    data := []values{
        {[]byte("abc"), []byte("de"),},
        {[]byte("abc"), []byte("ijk"),},
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
