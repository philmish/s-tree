package typed

type TypedRadix struct {
	*TypedTree
}

func newRadix() *TypedRadix {
	return &TypedRadix{
		newTree(),
	}
}

func (tr *TypedRadix) AddSequence(seq []interface{}) error {
	if tr.IsEmpty() {
		n, err := NewTypedNode("ROOT", nil)
		if err != nil {
			return err
		}
		tr.Root = n
	}
	if tr.Depth() == 0 {
		err := tr.AddBranch(seq)
		return err
	}
	lvlCursor := 0
	nodeCursor := tr.Root
	for _, val := range seq {
		if tr.Depth()-1 < lvlCursor {
			tr.Levels = append(tr.Levels, newLevel())
		}
		exists, err := tr.Levels[lvlCursor].findNodeWithValue(val)
		if err != nil {
			return err
		}
		if exists != nil {
			nodeCursor = exists
			lvlCursor++
			continue
		}
		n, err := NewTypedNode(val, nodeCursor)
		if err != nil {
			return err
		}
		nodeCursor = n
		tr.Levels[lvlCursor].addNode(n)
		lvlCursor++
	}
	return nil
}
