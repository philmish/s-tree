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

func (lvl *TypedTreeLevel) findNodeWithValue(val interface{}) (*TypedNode, error) {
	cursor, err := NewTypedNode(val, nil)
	if err != nil {
		return nil, err
	}
	for _, n := range lvl.nodes {
		v, err := n.GetValue()
		if err != nil {
			return nil, err
		}
		if n.Type == cursor.Type && val == v {
			return n, nil
		}
	}
	return nil, nil
}
