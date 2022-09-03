package pkg

type TreeLevel struct {
	Nodes []*Node
}

func newLevel() *TreeLevel {
	nodes := make([]*Node, 0)
	return &TreeLevel{
		Nodes: nodes,
	}
}

func (tl *TreeLevel) Append(n *Node) error {
    tl.Nodes = append(tl.Nodes, n)
    return nil
}

func (tl TreeLevel) Length() int {
	return len(tl.Nodes)
}

func (tl *TreeLevel) GetNodeByValue(value []byte) *Node {
	for _, n := range tl.Nodes {
		if n.CompareVal(value) {
			return n
		}
	}
	return nil
}
