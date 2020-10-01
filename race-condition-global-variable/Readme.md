### Demo tình trạng race condition giữa 2 goroutine khi cùng truy cập vào 1 biến global

#### Có 2 khả năng xảy ra:

1. TH1: Viết trước, đọc sau

![Viết trước, đọc sau](img/race-1.png?raw=true "Viết trước, đọc sau")

2. TH2: Đọc trước, viết sau

![Đọc trước, viết sau](img/race-2.png?raw=true "Đọc trước, viết sau")