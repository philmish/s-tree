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
}
