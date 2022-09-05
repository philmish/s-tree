package kvdb

import "testing"

func TestClientPing(t *testing.T) {
	server := NewDB("/tmp/kvdbtest")
	server.Run()
	defer server.Stop()
	go server.serve()
	client := KvClient{
		Addr: "/tmp/kvdbtest",
	}
	err := client.Ping()
	if err != nil {
		t.Errorf("Failed to ping server: %s\n", err.Error())
	}
}

func TestClientSet(t *testing.T) {
	server := NewDB("/tmp/kvdbtest")
	server.Run()
	defer server.Stop()
	go server.serve()
	client := KvClient{
		Addr: "/tmp/kvdbtest",
	}
	err := client.Set("a", "b")
	if err != nil {
		t.Errorf("Failed to ping server: %s\n", err.Error())
	}
}
