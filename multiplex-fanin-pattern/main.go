package main

import (
	"fmt"
	"time"
)

func main() {
	c := returnChannel("hello", 100)
	d := returnChannel("ahihi", 50)

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
	// 2 anonymous goroutine (ở dòng số 35) thay phiên nhau
	// gửi message vào channel result mà không phải chờ
	go func(){
		// <-chan1 là để móc message từ chan1 ra
		// khi đó <- <-chan1 sẽ là message từ chan1 được gửi vào channel result
		for {
			result <- <-chan1
		}
		
	}()

	go func(){
		// <-chan2 là để móc message từ chan2 ra
		// khi đó <- <-chan2 sẽ là message từ chan2 được gửi vào channel result
		for {
			result <- <-chan2
		}
		
	}()

	return result
}