package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

// Pool là một kho chứa đồ cho các goroutine xài chung
// Số lượng đồ xài chung đc biểu diễn dưới dạng capacity của 1 buffered channel
type Pool struct {
	// Kiểm tra xem kho còn mở không
	isClose bool

	// Sản xuất đồ mới trong trường hợp kho không còn đồ
	factory func() (io.Closer, error)

	// Các đồ trong kho
	items chan io.Closer

	// Tránh race condition khi các goroutine vào xem kho
	m sync.Mutex
}

// Hàm New để xây 1 kho chứa đồ mới
func New(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Kích thước kho phải lớn hơn 0 chứ")
	}

	return &Pool{
		factory: fn,
		items:   make(chan io.Closer, size),
		isClose: false,
	}, nil
}

// Hàm Acquire để lấy đồ từ kho
// Nếu kho có đồ ( buffered channel có item ) thì lấy
// Còn nếu không có thì tự tạo mới đồ, dùng xong thì lại đưa vào kho bằng hàm Release
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case item, ok := <-p.items:
		if !ok {
			fmt.Println("Kho đóng rồi")
			return nil, errors.New("Kho đóng mất rồi")
		}
		fmt.Println("Lấy được đồ rồi", item)
		return item, nil
	default:
		fmt.Println("Kho hết đồ rồi. Kho sẽ tạo đồ mới")
		return p.factory()
	}
}

// Hàm Release để trả đồ về kho
// Nếu kho đã đóng thì thôi, hủy món đồ đang định trả
// Nếu kho đã đầy đồ thì cũng thôi, hủy món đồ đang định trả đi
func (p *Pool) Release(item io.Closer) {
	p.m.Lock()

	if p.isClose {
		item.Close()
		return
	}

	select {
	case p.items <- item:
		fmt.Println("Trả đồ về kho")
	default:
		fmt.Println("Thôi kho full đồ rồi hủy món đồ đi thôi")
		item.Close()
	}

	defer p.m.Unlock()

}

// Hàm Close để đóng kho
func (p *Pool) Close() {
	p.m.Lock()

	if p.isClose {
		return
	}

	p.isClose = true
	close(p.items)

	for item := range p.items {
		item.Close()
	}

	defer p.m.Unlock()
}