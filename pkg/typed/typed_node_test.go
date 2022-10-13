package typed

import (
	"testing"
)

func TestStrNode(t *testing.T) {
	val := "Hello World"
	node, err := NewTypedNode(val, nil)
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
	node, err := NewTypedNode(val, nil)
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
	node, err := NewTypedNode(val, nil)
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

func TestIntSliceNode(t *testing.T) {
	val := []int{1}
	node, err := IntSliceNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create IntSliceNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of IntSliceNode", err.Error())
	}
	if ret.([]int)[0] != val[0] {
		t.Errorf("Expected %v found %v in value of IntSliceNode", val, ret)
	}
}

func TestBoolSliceNode(t *testing.T) {
	val := []bool{true}
	node, err := BoolSliceNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create BoolSliceNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of BoolSliceNode", err.Error())
	}
	if ret.([]bool)[0] != val[0] {
		t.Errorf("Expected %v found %v in value of BoolSliceNode", val, ret)
	}
}

func TestStrStrMapNode(t *testing.T) {
	val := map[string]string{
		"Hello": "World",
	}
	node, err := StrStrMapNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create StrStrMapNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of StrStrMapNode", err.Error())
	}
	if ret.(map[string]string)["Hello"] != val["Hello"] {
		t.Errorf("Expected %v found %v in value of StrStrMapNode", val, ret)
	}
}
