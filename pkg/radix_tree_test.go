package pkg

import (
	"testing"
)

func TestPush(t *testing.T) {
	tree := NewRadix()
	err := tree.PushStr("abcde")
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	if tree.Depth() != 5 {
		t.Errorf("Expected 5 levels, found %d\n", tree.Depth())
	}
}

func TestSearch(t *testing.T) {
	rt := NewRadix()
	err := rt.PushStr("abcde")
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	err = rt.SearchStr("abc")
	if err != nil {
		t.Errorf("Not found: %s\n", err.Error())
	}
}
