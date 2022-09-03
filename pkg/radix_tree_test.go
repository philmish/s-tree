package pkg

import (
	"testing"
)

func TestTreePush(t *testing.T) {
	tree := New()
	err := tree.Push("abcde")
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	if tree.t.Depth() != 5 {
		t.Errorf("Expected 5 levels, found %d\n", tree.t.Depth())
	}
}
