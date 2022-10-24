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

func TestIsEqualTo(t *testing.T) {
	val := "Hello World"
	node, err := NewTypedNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create first StrNode for is equal.", err.Error())
	}
	val2 := "Hello World"
	node2, err := NewTypedNode(val2, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create second StrNode for is equal.", err.Error())
	}
	if !node.IsEqualTo(node2) {
		t.Errorf("Expected nodes to be equal.")
	}
	val3 := "World"
	node3, err := NewTypedNode(val3, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create third StrNode for is equal.", err.Error())
	}
	if node.IsEqualTo(node3) || node2.IsEqualTo(node3) {
		t.Errorf("Expected nodes 1 and 2 to not be equal to node 3.")
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

func TestStrIntMapNode(t *testing.T) {
	val := map[string]int{
		"Hello": 1,
	}
	node, err := StrIntMapNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create StrIntMapNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of StrIntMapNode", err.Error())
	}
	if ret.(map[string]int)["Hello"] != val["Hello"] {
		t.Errorf("Expected %v found %v in value of StrIntMapNode", val, ret)
	}
}

func TestStrBoolMapNode(t *testing.T) {
	val := map[string]bool{
		"Hello": true,
	}
	node, err := StrBoolMapNode(val, nil)
	if err != nil {
		t.Errorf("%s\nFailed to create StrBoolMapNode", err.Error())
	}
	ret, err := node.GetValue()
	if err != nil {
		t.Errorf("%s\nFailed to read value of StrBoolMapNode", err.Error())
	}
	if ret.(map[string]bool)["Hello"] != val["Hello"] {
		t.Errorf("Expected %v found %v in value of StrBoolMapNode", val, ret)
	}
}
