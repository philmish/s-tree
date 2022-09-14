package kvdb

import (
	"net"
	"sync"
	"testing"
	"time"
)

func TestServerStartAndShutdown(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	time.Sleep(time.Second * 1)
	server.Stop()
}

func TestServerPing(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	defer server.Stop()
	buf := make([]byte, 256)
	time.Sleep(time.Second * 1)
	start := time.Now()
	conn, err := net.Dial("unix", "/tmp/TestKvDB")
	defer conn.Close()
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
	elapsed := time.Since(start)
	t.Logf("Ping took %s", elapsed)
}

func TestThreadedServerPing(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	defer server.Stop()
	buf := make([]byte, 256)
	time.Sleep(time.Second * 1)
	wg := new(sync.WaitGroup)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			conn, err := net.Dial("unix", "/tmp/TestKvDB")
			defer func() {
				conn.Close()
				wg.Done()
			}()
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
		}()
		wg.Wait()
	}
}

func TestServerSet(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	defer server.Stop()
	buf := make([]byte, 256)
	time.Sleep(time.Second * 1)
	conn, err := net.Dial("unix", "/tmp/TestKvDB")
	defer conn.Close()
	if err != nil {
		t.Errorf("error connecting to server: %s", err.Error())
	}
	n, err := conn.Write([]byte("SET a b"))
	if err != nil {
		t.Errorf("Failed to write to server")
	}
	n, err = conn.Read(buf)
	if err != nil {
		t.Errorf("Failed to read response from server: %s", err.Error())
	}
	msg := string(buf[:n])
	if msg != "RESULT SUCCESS" {
		t.Errorf("Expected RESULT SUCCESS recieved %s\n", msg)
	}
}

func TestServerGet(t *testing.T) {
	server := NewServer("/tmp/TestKvDB")
	defer server.Stop()
	buf := make([]byte, 256)
	time.Sleep(time.Second * 1)
	startWrite := time.Now()
	conn, err := net.Dial("unix", "/tmp/TestKvDB")
	defer conn.Close()
	if err != nil {
		t.Errorf("error connecting to server: %s", err.Error())
	}
	n, err := conn.Write([]byte("SET a b"))
	if err != nil {
		t.Errorf("Failed to write to server")
	}
	n, err = conn.Read(buf)
	if err != nil {
		t.Errorf("Failed to read response from server: %s", err.Error())
	}
	msg := string(buf[:n])
	if msg != "RESULT SUCCESS" {
		t.Errorf("Expected RESULT SUCCESS recieved %s\n", msg)
	}
	elapsedWrite := time.Since(startWrite)
	conn2, err := net.Dial("unix", "/tmp/TestKvDB")
	defer conn2.Close()
	startRead := time.Now()
	if err != nil {
		t.Errorf("error connecting to server for reading: %s", err.Error())
	}
	n, err = conn2.Write([]byte("GET a"))
	if err != nil {
		t.Errorf("Failed to write to server for Get")
	}
	buf = make([]byte, 256)
	n, err = conn2.Read(buf)
	if err != nil {
		t.Errorf("Failed to read from server for Get")
	}
	msg = string(buf[:n])
	if msg != "RESULT b" {
		t.Errorf("Expected RESULT b, recieved %s", msg)
	}
	elapsedRead := time.Since(startRead)
	elapsed := elapsedWrite + elapsedRead
	t.Logf("Write + Read took about %s", elapsed)
}
