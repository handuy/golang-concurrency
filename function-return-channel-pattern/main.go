package main

import "fmt"

func main() {
	c := returnChannel("hello")

	// <-c là một blocking operation
	// Khi nào bên anonymous routine ở dòng 27 gửi dữ liệu vào channel c thì 
	// code bên trong vòng lặp mới chạy
	// 5 lần chạy vòng lặp là 5 lần hàm main phải chờ anonymous routine gửi dữ liệu vào channel
	for i := 0; i < 5; i++ {
		fmt.Println("Hàm main nhận message:", <-c)
	}

	// Vòng lặp for chạy xong hết mới đến lượt chạy
	fmt.Println("Thoát khỏi hàm main")
}

func returnChannel(msg string) chan string {
	channel := make(chan string)

	// Hàm anonymous này chạy trên 1 goroutine riêng
	// Nếu số lần chạy vòng lặp ít hơn số lần chạy vòng lặp của hàm main (ở dòng số 12)
	// thì sẽ xảy ra hiện tượng deadlock
	// do bên hàm main channel nó vẫn chờ message được gửi sang từ phía này
	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintln("Gửi message", msg, "lần thứ", i)
		}
	}()

	// Hàm returnChannel vẫn chạy bình thường trên main goroutine
	// và trả về channel
	// Channel này đóng vai trò synchronize hoạt động của hàm main và anonymous goroutine ở dòng số 27
	fmt.Println("Thoát khỏi hàm returnChannel")
	return channel
}
