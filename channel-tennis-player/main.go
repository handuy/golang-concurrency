package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Quả bóng tennis sẽ được bắn vào channel này
	court := make(chan int)

	wg := &sync.WaitGroup{}
	// Có 2 player đang chơi tennis
	wg.Add(2)
	go player("Nadal", court, wg)
	go player("Djokovic", court, wg)

	// Bắt đầu trận đấu bằng cách gửi quả bóng vào channel
	// Ban đầu cả 2 player ( 2 goroutine ) đều đang ở trạng thái chờ có
	// quả bóng được gửi vào channel court ( dòng code số 35 )
	court <- 1
	wg.Wait()
}

func player(name string, court chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		ball, ok := <-court
		// Nếu channel court bị đóng --> player đối diện ( 1 goroutine khác ) đã đánh trượt
		// Thông báo player đã thắng
		if !ok {
			fmt.Println("Player", name, "won")
			return
		}

		// Thuật toán random để kiểm tra xem player khi nhận được bóng
		// từ channel court nhả ra ở dòng số 35
		// thì có đánh trúng quả bóng không
 		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println("Player", name, "đánh trượt quả bóng", ball, "mất rồi :(")
		
		 	// Thông báo cho player đối diện ( 1 goroutine khác ) là mình đã đánh trượt :(
			// bằng cách đóng channel court. Khi đó goroutine ở phía bên kia của channel court
			// nhận được tín hiệu đóng channel ( dòng code số 35, ok sẽ bằng false ) và 
			// cũng sẽ return thoát khỏi vòng lặp
			close(court)
			return
		}

		fmt.Println("Player", name, "trả giao bóng", ball)
		ball++

		// Bắn quả bóng vào channel để cho goroutine đầu bên kia nhận bóng
		court <- ball
	}
}
