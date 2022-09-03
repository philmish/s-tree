package pkg

import "strings"

type RadixTree struct {
    t *Tree
}

func New() *RadixTree {
    return &RadixTree{
        t: createTree(),
    }
}

func valsFromString(data string) values {
    symbols := strings.Split(data, "")
	vals := make(values, 0)
	var b []byte
	for len(symbols) > 0 {
		b, symbols = []byte(symbols[0]), symbols[1:]
		vals = append(vals, b)
	}
	return vals
}

func (rt *RadixTree) Push(data string) error {
    vals := valsFromString(data)
    return rt.t.PushSquence(vals)
}

func (rt *RadixTree) Search(data string) error {
    vals := valsFromString(data)
    return rt.t.SearchSequence(vals)
}
