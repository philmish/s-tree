package kvdb

import (
	"fmt"
	"net"
)

type DBClient struct {
	Addr string
}

func (client DBClient) Ping() (string, error) {
	conn, err := net.Dial("unix", client.Addr)
	defer conn.Close()
	if err != nil {
		return "", err
	}
	buf := make([]byte, 256)
	_, err = conn.Write([]byte("ping"))
	if err != nil {
		return "", err
	}
	n, err := conn.Read(buf)
	msg := buf[:n]
	return string(msg), nil
}

func (client DBClient) Set(key, val string) error {
	conn, err := net.Dial("unix", client.Addr)
	defer conn.Close()
	if err != nil {
		return err
	}
	buf := make([]byte, 16)
	stmt := fmt.Sprintf("SET %s %s", key, val)
	_, err = conn.Write([]byte(stmt))
	if err != nil {
		return err
	}
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	msg := buf[:n]
	if string(msg) != "RESULT SUCCESS" {
		return fmt.Errorf("Something went wrong setting key: %s\n", string(msg))
	}
	return nil
}
