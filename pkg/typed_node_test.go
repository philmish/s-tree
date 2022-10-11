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

func TestIntNode(t *testing.T) {
	val := 123
	node, err := IntNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create IntNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of IntNode", err.Error())
	}
	if ret != val {
		t.Errorf("Expected %d found %d in value of IntNode", val, ret)
	}
}
