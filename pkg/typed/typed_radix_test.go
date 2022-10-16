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
}
