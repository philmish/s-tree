package pkg

import (
	"bytes"
	"fmt"
)

type Node struct {
	Parent   *Node
	Children []*Node
	Value    value
}

func newNode(v value) *Node {
	return &Node{
		Parent:   nil,
		Children: make([]*Node, 0),
		Value:    v,
	}
}

func (n *Node) setParent(parent *Node) error {
	if &n == &parent {
		return fmt.Errorf("Node can not be its own parent")
	}
	n.Parent = parent
	return nil
}

func (n *Node) addChild(child *Node) error {
	if &n == &child {
		return fmt.Errorf("Node can not be its own parent")
	}
	err := child.setParent(n)
	if err != nil {
		return err
	}
	n.Children = append(n.Children, child)
	return nil
}

func CreateNode(v value, parent *Node) *Node {
	n := Node{
		Parent:   parent,
		Children: make([]*Node, 0),
		Value:    v,
	}
	parent.addChild(&n)
	return &n
}

func (n Node) CompareVal(data value) bool {
	res := bytes.Compare(data, n.Value)
	if res == 0 {
		return true
	} else {
		return false
	}
}
