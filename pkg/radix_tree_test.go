package pkg

import (
	"testing"
)

func TestPush(t *testing.T) {
	tree := New()
	err := tree.Push("abcde")
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	if tree.t.Depth() != 5 {
		t.Errorf("Expected 5 levels, found %d\n", tree.t.Depth())
	}
}

func TestSearch(t *testing.T) {
    rt := New()
	err := rt.Push("abcde")
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
    err = rt.Search("abc")
    if err != nil {
        t.Errorf("Not found: %s\n", err.Error())
    }
}
