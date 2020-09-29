package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		// Đăng kí 2 go routine đang chạy với wait group
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryCache(id); ok {
				fmt.Println("Hit cache !!!", b)
			}
			// Báo với wait group là chạy xông rồi
			wg.Done()
		}(id, wg)

		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("Hit db :(", b)
			}
			// Báo với wait group là chạy xông rồi
			wg.Done()
		}(id, wg)

	}

	// Chờ cho đến khi tất cả 20 go routine chạy xong đã 
	// rồi mới kết thúc hàm main
	wg.Wait()
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)

	for _, b := range Db {
		if b.ID == id {
			// Chỗ này gây ra race condition
			// Nhiều go routine cùng viết vào 1 cache[id]
			// Hoặc go routine đang viết trong khi có 1 go routine khác đang đọc
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false
}
