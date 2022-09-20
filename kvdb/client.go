package kvdb

import (
	"fmt"
	"net"
	"strings"
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

func (client DBClient) Get(key string) (string, error) {
	conn, err := net.Dial("unix", client.Addr)
	defer conn.Close()
	if err != nil {
		return "", err
	}
	stmt := fmt.Sprintf("GET %s", key)
	_, err = conn.Write([]byte(stmt))
	if err != nil {
		return "", err
	}
	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	msg := string(buf[:n])
	if !strings.HasPrefix(msg, "RESULT") {
		return "", fmt.Errorf("Error getting %s: %s\n", key, msg)
	}
	return strings.Split(msg, " ")[1], nil
}

func (client DBClient) Keys() ([]string, error) {
	conn, err := net.Dial("unix", client.Addr)
	defer conn.Close()
	if err != nil {
		return []string{}, err
	}
	_, err = conn.Write([]byte("KEYS"))
	if err != nil {
		return []string{}, err
	}
	buf := make([]byte, 1048)
	n, err := conn.Read(buf)
	if err != nil {
		return []string{}, err
	}
	res := buf[:n]
	data := strings.Split(string(res), " ")[1]
	return strings.Split(data, ","), nil

}
