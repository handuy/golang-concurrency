package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {

	// Mama tổng quản có nhiệm vụ theo đợi các goroutine chạy xong
	wg := &sync.WaitGroup{}

	// Có 2 thằng goroutine đang chạy đấy nhé. Nhớ theo dõi
	// 2 goroutine đang chạy đua với nhau
	// thằng add thì cố viết vào biến counter
	// thằng read thì cố đọc từ biến counter
	// Có lúc thằng add viết vào trước, rồi thằng read mới đến lượt đọc --> in ra 20
	// Có lúc thằng read lại được đọc trước, rồi thằng add mới viết --> in ra 0
	wg.Add(2)
	go add(20, wg)
	go read(wg)

	wg.Wait()
}

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("counter", counter)
}

func add(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	counter += amount
}
