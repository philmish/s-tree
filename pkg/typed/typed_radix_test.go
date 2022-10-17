package typed

import (
	"testing"
)

func TestAddSequence(t *testing.T) {
	tr := newRadix()
	seq := []interface{}{
		"Hello",
		1,
		"World",
		true,
	}
	err := tr.AddSequence(seq)
	if err != nil {
		t.Errorf("%s\nFailed to add first sequence to typed radix", err.Error())
	}
	seq2 := []interface{}{
		"Other",
		1,
		true,
	}
	err = tr.AddSequence(seq2)
	if err != nil {
		t.Errorf("%s\nFailed to add second sequence to typed radix", err.Error())
	}
	if tr.Depth() != 4 {
		t.Errorf("Expected radix tree to have Depth of 4, found %d", tr.Depth())
	}
	if tr.Levels[1].Length() != 1 {
		t.Errorf("Expected radix tree level 1 to contain 1 node, found %d", tr.Levels[1].Length())
	}
}
