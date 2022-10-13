package typed

type TypedTreeLevel struct {
	nodes []*TypedNode
}

func newLevel() *TypedTreeLevel {
	return &TypedTreeLevel{
		nodes: make([]*TypedNode, 0),
	}
}

func (tl *TypedTreeLevel) addNode(n *TypedNode) {
	tl.nodes = append(tl.nodes, n)
}

func (l *TypedTreeLevel) getLeafs() []*TypedNode {
	var res []*TypedNode
	for _, n := range l.nodes {
		if len(n.Children) == 0 {
			res = append(res, n)
		}
	}
	return res
}

func (l *TypedTreeLevel) getBranch(index int) *TypedNode {
	if index > len(l.nodes)-1 {
		return nil
	}
	r := l.nodes[index]
	return r
}
