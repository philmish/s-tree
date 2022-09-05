package kvdb

import (
	"fmt"
	"net"
)

type KvClient struct {
	Addr string
}

func (kvc *KvClient) send(query string, conn net.Conn) (string, error) {
	n, err := conn.Write([]byte(query))
	if err != nil {
		return "", err
	}
	buf := make([]byte, 256)
	n, err = conn.Read(buf)
	if err != nil {
		return "", err
	}
	str := buf[:n]
    return string(str), nil
}

func (kvc *KvClient) Ping() error {
	conn, err := net.Dial("unix", kvc.Addr)
	if err != nil {
		return err
	}
	defer conn.Close()
    result, err := kvc.send("PING", conn) 
    if err != nil {
        return err
    }
	if result != "PONG" {
		return fmt.Errorf("Expected PONG recieved: %s\n", result)
	}
	fmt.Println("RECIEVED: PONG")
	return nil
}

func (kvc *KvClient) Set(key, val string) error {
	conn, err := net.Dial("unix", kvc.Addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	query := fmt.Sprintf("SET %s %s", key, val)
    result, err := kvc.send(query, conn)
    if err != nil {
        return err
    }
	if result != "RES:SUCCESS" {
		return fmt.Errorf("Failed to set key/value pair: %s\n", result)
	}
	return nil
}
