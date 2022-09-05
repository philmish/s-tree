package kvdb

import (
	"os"
	"testing"
)

func TestDBStartAndStop(t *testing.T) {
    if err := os.RemoveAll("/tmp/kvdbtest"); err != nil {
        t.Error("Failed to kill socket.\n")
    }
	db := NewDB("/tmp/kvdbtest")
    db.Run()
    go db.serve()
    db.Stop()
}
