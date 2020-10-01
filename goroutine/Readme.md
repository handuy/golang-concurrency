### Minh họa quá trình chạy 2 goroutine

1. Giả sử goroutine printUpcase được Go runtime scheduler xếp chạy trước:

![Goroutine printUpcase được xếp chạy trước](img/routine-1.png?raw=true "Goroutine printUpcase được xếp chạy trước, còn goroutine printLowcase ở trong hàng chờ")

2. Đến lượt goroutine printLowcase được Go runtime scheduler xếp chạy:

![Đến lượt goroutine printLowcase được Go runtime scheduler xếp chạy](img/routine-2.png?raw=true "Đến lượt chạy của goroutine printLowcase")

#### Lưu ý: Thứ tự chạy sẽ khác nhau ở mỗi lần chạy code

Lúc thì printUpcase chạy hết 1 lượt rồi mới đến printLowcase, lúc thì printLowcase chạy được một nửa thì tạm dừng để cho printUpcase chạy, ... tùy vào Go runtime scheduler xếp lịch chạy như nào

![Thứ tự chạy khác nhau ở mỗi lần chạy code](img/routine-3.png?raw=true "Thứ tự chạy khác nhau ở mỗi lần chạy code")