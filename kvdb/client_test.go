package kvdb

import (
	"testing"
)

func TestClientPing(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	defer server.Stop()

	client := DBClient{Addr: "/tmp/TestKvDB"}
	msg, err := client.Ping()
	if err != nil {
		t.Errorf("Error pinging server: %s\n", err)
	}
	if msg != "PONG" {
		t.Errorf("Expected PONG, got %s\n", msg)
	}
}

func TestClientSet(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	defer server.Stop()

	client := DBClient{Addr: "/tmp/TestKvDB"}
	err := client.Set("test", "test")
	if err != nil {
		t.Errorf("Error seting key and value: %s\n", err)
	}
}
