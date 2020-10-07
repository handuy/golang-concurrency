package main

import (
	"fmt"
	"io"
	pool "learn-concurrency/pooling-pattern/pool"
	"sync"
	"time"

	"github.com/rs/xid"
)

type dbConn struct {
	connID string
}

func (db *dbConn) Close() error {
	fmt.Println("Đóng kết nối db")
	return nil
}

func createDbConn() (io.Closer, error) {
	newConnID := xid.New().String()
	fmt.Println("Tạo mới 1 DB connection")
	return &dbConn{
		connID: newConnID,
	}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(20)

	p, err := pool.New(createDbConn, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for query := 0; query <= 20; query++ {
		go func(q int) {
			runQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	fmt.Println("Finish")
	p.Close()
}

func runQuery(query int, p *pool.Pool) {
	newConn, err := p.Acquire()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer p.Release(newConn)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Query số", query, "từ connection ID", newConn.(*dbConn).connID)
}
