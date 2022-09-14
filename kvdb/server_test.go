package kvdb

import (
	"net"
	"testing"
	"time"
)

func TestServerStartAndShutdown(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	time.Sleep(time.Second * 3)
	server.Stop()
}

func TestServerPing(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	defer server.Stop()
	buf := make([]byte, 256)
	time.Sleep(time.Second * 1)
	conn, err := net.Dial("unix", "/tmp/TestKvDB")
	if err != nil {
		t.Errorf("error connecting to server: %s", err.Error())
	}
	n, err := conn.Write([]byte("ping"))
	if err != nil {
		t.Errorf("error reading from server: %s", err.Error())
	}
	n, err = conn.Read(buf)
	if err != nil {
		t.Errorf("Failed to read response from server: %s", err.Error())
	}
	msg := string(buf[:n])
	if msg != "PONG" {
		t.Errorf("Expected PONG, recieved: %s", msg)
	}
	conn.Close()
}
