package main

import (
	"fmt"
	"time"
)

func main() {
	c := returnChannel("hello", 2000)
	d := returnChannel("ahihi", 500)

	// Tin nhắn từ 2 channel c, d được bắn vào trong
	// cùng 1 channel là channelFromFanIn
	channelFromFanIn := fanIn(c, d)

	// <-channelFromFanIn là một blocking operation
	for i := 0; i < 10; i++ {
		// Mỗi lần chạy vòng lặp, <-channelFromFanIn sẽ chờ message được gửi đến
		// từ 1 trong 2 channel c, d
		// thằng nào gửi trước thì <-channelFromFanIn sẽ nhận và hiển thị lên hàm main 
		// rồi sau đó chạy sang vòng lặp tiếp
		fmt.Println("Hàm main nhận message từ channel", <-channelFromFanIn)
		fmt.Println("Xong lần chạy thứ", i)
	}

	// Vòng lặp for chạy xong hết mới đến lượt chạy
	fmt.Println("Thoát khỏi hàm main")
}

func returnChannel(msg string, sleepTime int) chan string {
	channel := make(chan string)

	// Hàm anonymous này chạy trên 1 goroutine riêng
	// Nếu số lần chạy vòng lặp ít hơn số lần chạy vòng lặp của hàm main (ở dòng số 13)
	// thì sẽ xảy ra hiện tượng deadlock
	// do bên hàm main channel nó vẫn chờ message được gửi sang từ phía này
	go func() {
		for i := 0; ; i++ {
			time.Sleep( time.Duration(sleepTime) * time.Millisecond )
			channel <- fmt.Sprintln("Gửi message", msg, "lần thứ", i)
		}
	}()

	fmt.Println("Thoát khỏi hàm returnChannel")
	return channel
}

func fanIn(chan1, chan2 chan string) chan string {
	result := make(chan string)
	// Chỉ cần 1 anonymous goroutine (so với 2 ở ví dụ multiplex-fanin-pattern)
	// để hứng kết quả từ 2 goroutine ở dòng số 37
	go func(){
		// Channel nào gửi trước thì thông tin sẽ được gửi vào channel result
		// để từ đó bắn vào hàm main (dòng số 22)
		for {
			select {
			case msg := <-chan1:
				result <- msg
			case msg := <-chan2:
				result <- msg
			}
		}
		
	}()

	return result
}