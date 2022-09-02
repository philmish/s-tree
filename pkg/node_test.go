package pkg

import (
	"testing"
)

type compareTestCase struct {
    valA string
    valB string
    expect bool
}

func (c compareTestCase)passes() bool {
    n1 := newNode([]byte(c.valA))
    equal := n1.CompareVal([]byte(c.valB))
    return c.expect == equal
}

func TestNodeCompare(t *testing.T) {
    data := []compareTestCase{
        {"a", "a", true},
        {"a", "b", false},
    }
    for _, tCase := range data {
        if !tCase.passes() {
            t.Errorf(
                "Failed compare for %s to %s equality to be %v",
                tCase.valB,
                tCase.valA,
                tCase.expect,
            )
        }
    }
}
