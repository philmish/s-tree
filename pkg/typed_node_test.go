package pkg

import (
	"testing"
)

func TestStrNode(t *testing.T) {
	val := "Hello World"
	node, err := StrNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create StrNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of StrNode", err.Error())
	}
	if ret != val {
		t.Errorf("Expected %s found %s in value of StrNode", val, ret)
	}
}
