package kvdb

import (
	"log"
	"net"
	"os"
	"sync"

	"github.com/philmish/s-tree/pkg"
)

type DBServer struct {
	t        *pkg.Tree
	addr     string
	listener net.Listener
	quit     chan interface{}
	wg       sync.WaitGroup
}

func NewServer(addr string) *DBServer {
	t := pkg.NewTree("")
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

func pingServer(c net.Conn) {
	defer c.Close()
	log.Printf("Connection: [%s]", c.RemoteAddr().Network())
	msg := []byte("PONG")
	_, err := c.Write(msg)
	if err != nil {
		log.Fatalf("error reading ping message")
	}
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
				pingServer(conn)
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
