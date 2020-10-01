package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Tổng cộng có 8 logical processor
	// fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(4)

	// Mama tổng quản có nhiệm vụ theo đợi các goroutine chạy xong
	wg := &sync.WaitGroup{}

	// Có 2 thằng goroutine đang chạy đấy nhé. Nhớ theo dõi
	wg.Add(2)

	go printPrime("A", wg)
	go printPrime("B", wg)

	fmt.Println("Run main")

	// Đợi cho đến khi 2 goroutine return thì mới kết thúc hàm main
	wg.Wait()
}

func printPrime(routineID string, wg *sync.WaitGroup) {
	// Báo với mama tổng quản là chạy xong rồi
	defer wg.Done()

	for outer := 2; outer < 1000000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue
			}
		}
		fmt.Println("From routine ", routineID, ": ", outer)
	}

	fmt.Println("Routine", routineID, " Completed")
}
