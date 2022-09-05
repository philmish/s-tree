package kvdb

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type KvDB struct {
	Addr     string
	DB       *KvStore
	listener net.Listener
	quit     chan interface{}
	wg       sync.WaitGroup
}

func NewDB(addr string) *KvDB {
	return &KvDB{
		Addr: addr,
		DB:   NewStore(),
		quit: make(chan interface{}),
		wg:   sync.WaitGroup{},
	}
}

func (kv *KvDB) handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	leng, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading: %#v\n", err)
		return
	}
	msg := buf[:leng]
	log.Printf("Recieved: %s\n", string(msg))
	parts := strings.Split(string(msg), " ")
	wg := sync.WaitGroup{}
	switch parts[0] {
	case "GET":
		query := parts[1]
		wg.Add(1)
		str, err := kv.DB.Get(query, &wg)
		wg.Wait()
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("ERR:%s", err.Error())))
			return
		}
		conn.Write([]byte(fmt.Sprintf("RES:%s", str)))
		return
	case "SET":
		if len(parts) < 3 {
			conn.Write([]byte("Invalid query length"))
			return
		}
		key := parts[1]
		val := parts[2]
		wg.Add(1)
		err := kv.DB.Set(key, val, &wg)
		wg.Wait()
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("ERR:%s", err.Error())))
			return
		}
		conn.Write([]byte("RES:SUCCESS"))
		return
    case "PING":
        conn.Write([]byte("PONG"))
        return
	default:
		conn.Write([]byte("ERR:UNKNOWN"))
		return
	}
}

func (kv *KvDB) serve() {
	conn, err := kv.listener.Accept()
	if err != nil {
		select {
		case <-kv.quit:
			return
		default:
			log.Println("Accept error: ", err)
		}
	} else {
		kv.wg.Add(1)
		go func() {
			kv.handleConn(conn)
			kv.wg.Done()
		}()
	}
}

func (kv *KvDB) Run() {
	l, err := net.Listen("unix", kv.Addr)
	if err != nil {
		log.Fatalf("Failed to start server.\nErr: %s\n", err.Error())
	}
	kv.listener = l
}

func (kv *KvDB) Stop() {
	close(kv.quit)
	kv.listener.Close()
	os.RemoveAll(kv.Addr)
}
