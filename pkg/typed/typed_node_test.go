package typed

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

func TestBoolNode(t *testing.T) {
	val := true
	node, err := BoolNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create BoolNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of BoolNode", err.Error())
	}
	if ret != val {
		t.Errorf("Expected %v found %v in value of BoolNode", val, ret)
	}
}

func TestStrSliceNode(t *testing.T) {
	val := []string{"Hello World"}
	node, err := StrSliceNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create StrSliceNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of StrSliceNode", err.Error())
	}
	if ret.([]string)[0] != val[0] {
		t.Errorf("Expected %v found %v in value of StrSliceNode", val, ret)
	}
}
