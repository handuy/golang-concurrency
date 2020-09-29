```go
for i := 0; i < 10; i++ {
    id := rnd.Intn(10) + 1

    go func(id int) {
        if b, ok := queryCache(id); ok {
            fmt.Println("Hit cache !!!", b)
        }
    }(id)

    go func(id int) {
        if b, ok := queryDatabase(id); ok {
            fmt.Println("Hit db :(", b)
        }
    }(id)
    
    fmt.Println("Book not found")
}
```

Chạy vòng lặp 10 lần:
- Mỗi lần sinh ra 2 go routine: Routine 1 chọc vào cache, routine 2 chọc vào db
- Như vậy tổng cộng sinh ra 10 x 2 = 20 routine
- Tuy nhiên chương trình sẽ chỉ chạy ra 10 dòng Book not found
- Nguyên nhân: Sau khi in ra dòng Book not found thứ 10, hàm main kết thúc trước khi có bất kỳ routine trả về kết quả