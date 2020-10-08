package main

import (
	"fmt"
	"time"
)

func main() {
	c := returnChannel("hello", 2000)
	for i := 0; i <= 30 ; i++ {
		// Mỗi lần chạy vòng lặp sẽ đánh giá 2 channel: c và time.After
		// xem thằng nào gửi message trước
		// thằng c thì cứ 2s mới lại gửi message
		// trong khi thằng time.After cứ 100ms lại gửi message
		// do đó 19 lần chạy đầu đều là message từ thằng c
		// đến lần thứ 20 mới đến lượt thằng c gửi message
		// 10 lần cuối đều là thằng time.After
		select {
		case msg := <- c:
			fmt.Println("Hồi âm từ channel:", msg)
		case <- time.After(time.Duration(100) * time.Millisecond):
			fmt.Println("Mãi không thấy phản hồi")
		}
	}

	// Vòng lặp for chạy xong hết mới đến lượt chạy
	fmt.Println("Thoát khỏi hàm main")
}

func returnChannel(msg string, sleepTime int) chan string {
	channel := make(chan string)

	// Hàm anonymous này chạy trên 1 goroutine riêng
	go func() {
		for i := 0; i <= 5 ; i++ {
			time.Sleep( time.Duration(sleepTime) * time.Millisecond )
			channel <- fmt.Sprintln("Gửi message", msg, "lần thứ", i)
		}
	}()

	// Hàm returnChannel vẫn chạy bình thường trên main goroutine
	// và trả về channel
	fmt.Println("Thoát khỏi hàm returnChannel")
	return channel
}
