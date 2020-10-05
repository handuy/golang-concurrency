package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// Kết nối hàm main với goroutine chạy task
	complete chan error

	// Kết nối hàm main với goroutine chạy task
	timeOut <-chan time.Time

	// Kết nối goroutine chạy task với OS
	interrupt chan os.Signal

	// Các task được chạy tuần tự trong goroutine
	tasks []func(id int)
}

func New(d time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeOut:   time.After(d),
		interrupt: make(chan os.Signal, 1),
	}
}

func (r *Runner) AddTask(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Chạy các task trên 1 goroutine khác và lắng nghe kết quả trên 2 channel
// runner.complete và runner.timeOut
func (r *Runner) Start() error {
	// Kết nối OS với channel r.interrupt
	// Nếu có sự kiện từ phía OS thì OS sẽ gửi tín hiệu đến channel r.interrupt
	// khi đó đoạn code ở dòng số 77 sẽ được chạy
	signal.Notify(r.interrupt, os.Interrupt)

	// Chạy các task trong r.tasks ở goroutine r.run()
	// và đặt channel r.complete để hứng kết quả từ goroutine này 
	go func(){
		r.complete <- r.run()
	}()

	select {
	// Nếu các task chạy trong khoảng thời gian quy định
	case err := <- r.complete:
		return err
	// Nếu bị timeout
	case <- r.timeOut:
		return errors.New("Timeout mất rồi")
	}
}

// Goroutine chạy task
// Trước khi chạy mỗi task thì lắng nghe tín hiệu từ OS thông qua channel runner.interrupt
func (r *Runner) run() error {
	for i := 0; i < len(r.tasks); i++ {
		if ok := r.checkInterrupt(); ok {
			return errors.New("Bị shutdown từ phía OS")
		}
		r.tasks[i](i)
	}
	return nil
}

// Lắng nghe tín hiệu OS gửi đến channel runner.interrupt
func (r *Runner) checkInterrupt() bool {
	select {
	// Nếu có tín hiệu từ phía OS (ví dụ người dùng bấm Ctrl + C) thì báo cho 
	// hàm run biết để dừng không chạy nữa
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	// Nếu OS không có tín hiệu gì thì báo cho hàm run tiếp tục chạy task
	default:
		return false
	}
}
