package main

import (
	"fmt"
	"sync"
)

func main() {
	// Mama tổng quản có nhiệm vụ theo đợi các goroutine chạy xong
	wg := &sync.WaitGroup{}

	// Có 2 thằng goroutine đang chạy đấy nhé. Nhớ theo dõi
	wg.Add(2)

	// 2 goroutines được tạo ra gắn với 1 logical processor
	// Go runtime scheduler liên tục switching 2 goroutine
	// Lúc thì printLowcase được chạy bởi logical processor
	// Lúc thì là printUpcase
	go printLowcase(wg)
	go printUpcase(wg)

	fmt.Println("Run main")

	// Đợi cho đến khi 2 goroutine return thì mới kết thúc hàm main
	wg.Wait()
}

func printLowcase(wg *sync.WaitGroup) {
	// Báo với mama tổng quản là chạy xong rồi
	defer wg.Done()
	
	for char := 'a'; char < 'a'+26; char++ {
		fmt.Printf("%c ", char)
	}
}

func printUpcase(wg *sync.WaitGroup) {
	// Báo với mama tổng quản là chạy xong rồi
	defer wg.Done()
	
	for char := 'A'; char < 'A'+26; char++ {
		fmt.Printf("%c ", char)
	}
}
