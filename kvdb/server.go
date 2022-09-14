package kvdb

import (
	"log"
	"net"
	"os"
	"sync"

	"github.com/philmish/s-tree/pkg"
)

type DBServer struct {
	t        *pkg.RadixTree
	addr     string
	listener net.Listener
	quit     chan interface{}
	wg       sync.WaitGroup
}

func NewServer(addr string) *DBServer {
	t := pkg.NewRadix()
	db := &DBServer{
		addr: addr,
		t:    t,
		quit: make(chan interface{}),
	}
	l, err := net.Listen("unix", db.addr)
	if err != nil {
		log.Fatalf("error listening: %s\n", err.Error())
	}
	db.listener = l
	db.wg.Add(1)
	go db.serve()
	return db
}

func (db *DBServer) handlConn(c net.Conn) {
	defer c.Close()
	log.Printf("Connection: [%s]", c.RemoteAddr().Network())
	buf := make([]byte, 256)
	n, err := c.Read(buf)
	if err != nil {
		log.Println("Failed to read from connection")
		return
	}
	cmd, err := parseCommand(buf[:n])
	if err != nil {
		log.Printf("Failed to parse command: %s\n", string(buf[:n]))
		return
	}
	resp := cmd.execute(db.t)
	_, err = c.Write([]byte(resp))
	if err != nil {
		log.Printf("Failed to send response: %s", resp)
		return
	}
	return
}

func (db *DBServer) serve() {
	defer db.wg.Done()

	for {
		conn, err := db.listener.Accept()
		if err != nil {
			select {
			case <-db.quit:
				return
			default:
				log.Println("Error accepting conn: ", err)
			}
		} else {
			db.wg.Add(1)
			go func() {
				db.handlConn(conn)
				db.wg.Done()
			}()
		}
	}
}

func (db *DBServer) Stop() {
	close(db.quit)
	db.listener.Close()
	db.wg.Wait()
	if err := os.RemoveAll(db.addr); err != nil {
		log.Fatal(err)
	}
}
