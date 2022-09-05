package kvdb

import (
	"net"
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

func TestDBSet(t *testing.T) {
	db := NewDB("/tmp/kvdbtest")
	db.Run()
	defer db.Stop()
	go db.serve()
	conn, err := net.Dial("unix", "/tmp/kvdbtest")
	defer conn.Close()
	if err != nil {
		t.Errorf("Failed to connect to DB")
	}
	query := "SET a b"
	n, err := conn.Write([]byte(query))
	if err != nil {
		t.Errorf("Failed to write to DB")
	}
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		t.Errorf("Failed to recieve from DB")
	}
	if n == 0 {
		t.Errorf("No answer recieved")
	}
	msg := buf[:n]
	t.Logf("Msg: %s\n", msg)
}
