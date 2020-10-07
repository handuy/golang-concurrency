package main

import (
	"fmt"
	"time"
)

func main() {
	c := returnChannel("hello", 2000)
	d := returnChannel("ahihi", 50)

	// <-c là một blocking operation
	// Khi nào bên anonymous routine ở dòng 35 gửi dữ liệu vào channel c thì
	// code bên trong vòng lặp mới chạy
	// 5 lần chạy vòng lặp là 5 lần hàm main phải chờ anonymous routine gửi dữ liệu vào channel
	for i := 0; i < 5; i++ {
		// Chờ channel c nhận message trước
		// rồi mới đến lượt channel d chờ
		fmt.Println("Hàm main nhận message từ channel c:", <-c)
		fmt.Println("Hàm main nhận message từ channel d:", <-d)
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

	// Hàm returnChannel vẫn chạy bình thường trên main goroutine
	// và trả về channel
	// Channel này đóng vai trò synchronize hoạt động của hàm main và 
	// anonymous goroutine ở dòng số 35
	fmt.Println("Thoát khỏi hàm returnChannel")
	return channel
}
