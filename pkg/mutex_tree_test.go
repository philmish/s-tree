package pkg

import (
	"fmt"
	"sync"
	"testing"
)

func TestPushMutexTree(t *testing.T) {
    tree := createTree()
    mt := MuTree{
        t: tree,
    }
    data := []values{
        {[]byte("abc"), []byte("de"),},
        {[]byte("abc"), []byte("ijk"),},
    }
    wg := sync.WaitGroup{}
    for _, i := range data {
        wg.Add(1)
        go func(n values, wg *sync.WaitGroup) {
            err := mt.Push(n, wg)
            if err != nil {
                fmt.Println("Err")
            }
        }(i, &wg)
    }
    wg.Wait()
    if mt.t.Depth() != 2 {
        t.Logf("%v\n", mt.t.Levels[0].Length())
        t.Errorf("Expected depth 2 got %d", mt.t.Depth())
    }

}
