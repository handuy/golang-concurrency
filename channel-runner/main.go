package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Gậy tiếp sức của các runner
	// Channel dùng để kết nối thông tin giữa các runner
	item := make(chan int)
	number := 1

	wg := &sync.WaitGroup{}
	// Sẽ có 4 runner ( 4 goroutine ) được khởi chạy
	// Tuy nhiên hàm main chỉ cần đợi goroutine cuối cùng chạy xong
	// Thì sẽ kết thúc, qua đó shutdown luôn 3 goroutine còn lại
	wg.Add(1)

	// Runner đầu tiên khởi chạy
	go runner(number, item, wg)

	// Thông báo với goroutine đầu tiên là cuộc chạy đã bắt đầu
	item <- 1

	// Blocking hàm main cho đến khi runner (goroutine) số 4 chạy xong
	wg.Wait()
}

func runner(name int, court chan int, wg *sync.WaitGroup) {
	// Đợi cho đến khi có goroutine (runner) trao gậy thì mới chạy tiếp
	myTurn := <-court

	fmt.Println("Runner số", name, "đang chạy")

	var newRunner int

	// Trong lúc runner (goroutine) hiện tại đang chạy
	// Một runner (goroutine) khác bước vào đường đua và sẵn sàng nhận gậy
	// Tuy nhiên do chỉ giới hạn 4 runner nên cần check biến name
	if myTurn != 4 {
		newRunner = name + 1
		fmt.Println("Runner số", newRunner, "bước vào đường chạy")

		// Runner (goroutine) mới này mới chỉ đứng trên đường chạy thôi chứ chưa chạy
		// Do đang bị blocking ở dòng code số 31
		// Chỉ khi runner (goroutine) name gửi gậy vào channel court thì
		// runner (goroutine) newRunner mới chạy tiếp
		go runner(newRunner, court, wg)
	}

	// Giả lập là runner (goroutine) name đang chạy 
	time.Sleep(100 * time.Millisecond)

	// Đến runner (goroutine) số 4 thì thôi ko chạy nữa
	// Thông báo cho wg là đã chạy xong và return
	// Lúc này hàm main sẽ được gỡ blocking ở dòng wg.Wait() và thoát
	if myTurn == 4 {
		fmt.Println("Runner số", myTurn, "đã hoàn thành vòng thi")
		wg.Done()
		return
	}
	

	fmt.Println("Runner số", name, "đã hoàn thành vòng đua và đang trao gậy cho runner số", newRunner)

	// Trao gậy cho runner (goroutine) 
	court <- newRunner
}
