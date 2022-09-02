package pkg

import (
	"bytes"
	"fmt"
)

type Node struct {
    Parent *Node
    Children []*Node
    Value []byte
}

func newNode(value []byte) *Node {
    return &Node{
        Parent: nil,
        Children: make([]*Node, 0),
        Value: value,
    }
}

func (n *Node)setParent(parent *Node) error {
    if &n == &parent {
        return fmt.Errorf("Node can not be its own parent")
    }
    n.Parent = parent
    return nil
}

func (n *Node)addChild(child *Node) error {
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

func CreateNode(value []byte, parent *Node) *Node {
    return &Node{
        Parent: parent,
        Children: make([]*Node, 0),
        Value: value,
    }
}

func (n Node)CompareVal(data []byte) bool {
    res := bytes.Compare(data, n.Value)
    if res == 0 {
        return true
    } else {
        return false
    }
}
