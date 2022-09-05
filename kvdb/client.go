package kvdb

import (
	"fmt"
	"net"
)

type KvClient struct {
	Addr string
}

func (kvc *KvClient) Ping() error {
	conn, err := net.Dial("unix", kvc.Addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	msg := []byte("PING")
	n, err := conn.Write(msg)
	if err != nil {
		return err
	}
	fmt.Println("SEND PONG")
	if n == 0 {
		return fmt.Errorf("No data send")
	}
	buf := make([]byte, 256)
	n, err = conn.Read(buf)
	if err != nil {
		return err
	}
	str := buf[:n]
	if string(str) != "PONG" {
		return fmt.Errorf("Expected PONG recieved: %s\n", string(str))
	}
	fmt.Println("RECIEVED: PONG")
	return nil
}
