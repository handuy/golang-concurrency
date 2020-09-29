```go
wg := &sync.WaitGroup{}

for i := 0; i < 10; i++ {
    id := rnd.Intn(10) + 1
    // Đăng kí 2 go routine đang chạy với wait group
    wg.Add(2)

    go func(id int, wg *sync.WaitGroup) {
        if b, ok := queryCache(id); ok {
            fmt.Println("Hit cache !!!", b)
        }
        // Báo với wait group là chạy xông rồi
        wg.Done()
    }(id, wg)

    go func(id int, wg *sync.WaitGroup) {
        if b, ok := queryDatabase(id); ok {
            fmt.Println("Hit db :(", b)
        }
        // Báo với wait group là chạy xông rồi
        wg.Done()
    }(id, wg)

}

// Chờ cho đến khi tất cả 20 go routine chạy xong đã 
// rồi mới kết thúc hàm main
wg.Wait()
```

Chạy vòng lặp 10 lần:
- Mỗi lần sinh ra 2 go routine: Routine 1 chọc vào cache, routine 2 chọc vào db
- Như vậy tổng cộng sinh ra 10 x 2 = 20 routine
- Tạo 1 Wait Group, như một tổng quản phụ trách theo dõi 20 go routine
- Trong mỗi vòng lặp, đăng kí 2 go routine với wait group để nó theo dõi. Mỗi khi có 1 go routine chạy xong thì báo với wait group là tôi chạy xong rồi, không phải chờ nữa
- Bên ngoài vòng lặp, đặt hàm wg.Wait() để hàm main chờ cả 20 go routine chạy xong hết thì mới quit

Version này khắc phục race conditon bằng cách xài mutex. Chi tiết xem trong comment code 