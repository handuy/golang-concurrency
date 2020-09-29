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
	mu := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		// Đăng kí 2 go routine đang chạy với wait group
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, mu *sync.Mutex) {
			if b, ok := queryCache(id, mu); ok {
				fmt.Println("Hit cache !!!", b)
			}
			// Báo với wait group là chạy xông rồi
			wg.Done()
		}(id, wg, mu)

		go func(id int, wg *sync.WaitGroup, mu *sync.Mutex) {
			if b, ok := queryDatabase(id, mu); ok {
				fmt.Println("Hit db :(", b)
			}
			// Báo với wait group là chạy xông rồi
			wg.Done()
		}(id, wg, mu)
		time.Sleep(1000 * time.Millisecond)
	}

	// Chờ cho đến khi tất cả 20 go routine chạy xong đã 
	// rồi mới kết thúc hàm main
	wg.Wait()
}

func queryCache(id int, mu *sync.Mutex) (Book, bool) {
	// Chỉ 1 routine được phép đọc từ cache
	// Đọc xong mới nhả ra cho thằng routine khác đọc/viết
	mu.Lock()
	b, ok := cache[id]
	mu.Unlock()
	return b, ok
}

func queryDatabase(id int, mu *sync.Mutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)

	for _, b := range Db {
		if b.ID == id {
			// Khắc phục race condition bằng mutex
			// Chỉ 1 go routine đc phép write vào cache tại 1 thời điểm
			// Write xong thì mới nhả ra cho routine khác viết hoặc đọc
			mu.Lock()
			cache[id] = b
			mu.Unlock()
			return b, true
		}
	}

	return Book{}, false
}
