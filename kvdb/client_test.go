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

func TestClientSetGet(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	defer server.Stop()

	client := DBClient{Addr: "/tmp/TestKvDB"}
	err := client.Set("test", "test")
	if err != nil {
		t.Errorf("Error seting key and value: %s\n", err)
	}
	val, err := client.Get("test")
	if err != nil {
		t.Errorf("Error getting value for test: %s\n", err)
	}
	if val != "test" {
		t.Errorf("Expected test, got %s\n", val)
	}
}
