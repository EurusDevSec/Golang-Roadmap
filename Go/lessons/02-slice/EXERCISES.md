# Bài tập 02 — Slice từ cơ bản đến sử dụng thực tế

Các bài 1–7 ôn nền tảng. Các bài 8–18 luyện cách dùng slice đúng trong code thực tế nhưng không có thuật toán khó.

## Bài 1 — Tạo slice

```go
services := []string{"api", "db", "cache"}
```

Output:

```text
[api db cache]
length=3
```

## Bài 2 — Thêm phần tử

Thêm `worker` vào slice trên.

```text
before: [api db cache]
after:  [api db cache worker]
```

Gợi ý: `services = append(services, "worker")`.

## Bài 3 — Sửa phần tử

Đổi `db` thành `postgres`.

```text
[api postgres cache]
```

## Bài 4 — Tính tổng

Cho `[]int{5, 10, 15, 20}`.

```text
sum=50
```

## Bài 5 — Lọc số chẵn

Cho `[]int{1, 2, 3, 4, 5, 6}`. Tạo slice kết quả mới.

```text
input:  [1 2 3 4 5 6]
output: [2 4 6]
```

Gợi ý: nếu số chẵn thì `append` vào `result`.

## Bài 6 — Tìm tên service

Cho `[]string{"api", "db", "cache"}` và `target := "db"`.

```text
found=true
```

Thử `target := "worker"`, output là `found=false`.

## Bài 7 — Array khác slice

Chạy hai khai báo:

```go
array := [3]int{1, 2, 3}
slice := []int{1, 2, 3}
slice = append(slice, 4)
```

Output:

```text
array=[1 2 3] length=3
slice=[1 2 3 4] length=4
```

Viết một câu giải thích vì sao không thể `append` vào array.

## Bài 8 — `make` với length

Tạo slice có sẵn 3 phần tử bằng:

```go
services := make([]string, 3)
```

Gán lần lượt `api`, `db`, `cache` bằng chỉ số.

```text
services=[api db cache]
len=3 cap=3
```

Mục tiêu: hiểu rằng `make([]T, length)` đã tạo sẵn các phần tử zero value.

## Bài 9 — `make` với capacity

Tạo slice rỗng nhưng chuẩn bị chỗ cho 3 phần tử:

```go
services := make([]string, 0, 3)
```

Dùng `append` thêm `api`, `db`, `cache`.

```text
before: [] len=0 cap=3
after:  [api db cache] len=3 cap=3
```

Không gán `services[0]` khi length đang bằng `0`, vì sẽ panic.

## Bài 10 — Thêm nhiều phần tử

Cho:

```go
services := []string{"api"}
more := []string{"db", "cache", "worker"}
```

Dùng một lần `append` để tạo kết quả:

```text
[api db cache worker]
```

Gợi ý: khi append một slice khác, cần toán tử `...`.

## Bài 11 — Cắt slice

Cho:

```go
ports := []int{22, 80, 443, 3000, 8080}
```

Lấy phần từ chỉ số `1` đến trước chỉ số `4`.

```text
original: [22 80 443 3000 8080]
part:     [80 443 3000]
```

Mục tiêu: hiểu cú pháp `slice[start:end]`, trong đó `end` không được lấy.

## Bài 12 — Quan sát backing array

Tiếp tục từ bài 11, đổi phần tử đầu tiên của `part` thành `81`.

Output cần quan sát:

```text
part:     [81 443 3000]
original: [22 81 443 3000 8080]
```

Viết một câu giải thích vì sao `original` cũng thay đổi.

## Bài 13 — Copy độc lập

Cho:

```go
original := []string{"dev", "staging", "prod"}
```

Tạo `cloned` bằng `make` và `copy`, sau đó đổi `cloned[0]` thành `local`.

```text
original: [dev staging prod]
cloned:   [local staging prod]
```

Yêu cầu: `original` không được thay đổi.

## Bài 14 — Sửa slice trong function

Hoàn thành function:

```go
func double(values []int) {
	// TODO
}
```

Input và output:

```text
before: [1 2 3]
after:  [2 4 6]
```

Không trả slice mới. Sửa từng phần tử bằng chỉ số.

## Bài 15 — Append trong function

Hoàn thành:

```go
func addService(services []string, name string) []string {
	// TODO
}
```

Sử dụng:

```go
services := []string{"api", "db"}
services = addService(services, "cache")
```

Output:

```text
[api db cache]
```

Best practice: function có `append` nên trả slice và phía gọi phải nhận lại kết quả.

## Bài 16 — Xóa một phần tử

Cho:

```go
services := []string{"api", "db", "cache", "worker"}
index := 1
```

Xóa phần tử tại `index`, giữ nguyên thứ tự.

```text
before: [api db cache worker]
after:  [api cache worker]
```

Gợi ý: ghép phần trước index với phần sau index bằng `append` và `...`.

Chỉ làm với index hợp lệ; chưa cần viết xử lý lỗi.

## Bài 17 — Nil slice và empty slice

Chạy và quan sát:

```go
var nilSlice []int
emptySlice := make([]int, 0)
```

In `len` và kết quả so sánh với `nil`, sau đó append `10` vào cả hai.

```text
nilSlice:   len=0 isNil=true
emptySlice: len=0 isNil=false
after append: [10] [10]
```

Kết luận cần viết: cả hai đều có thể dùng `len`, `range` và `append` an toàn.

## Bài 18 — Xóa toàn bộ giá trị bằng `clear`

Cho:

```go
values := []int{10, 20, 30}
```

Dùng `clear(values)`.

```text
before: [10 20 30]
after:  [0 0 0]
len=3
```

Mục tiêu: hiểu rằng `clear` đưa phần tử về zero value, không làm length thành `0`.

## Hoàn thành khi

Bạn đã nắm và dùng được slice khi:

- Tự làm được bài 1–10 mà không xem gợi ý.
- Phân biệt được `len` và `cap`.
- Giải thích được vì sao subslice có thể làm dữ liệu gốc thay đổi.
- Biết dùng `copy` khi cần dữ liệu độc lập.
- Biết function có `append` nên trả slice mới.
- Làm được bài 14–17.

Chưa cần làm `CHALLENGES.md` nếu các ý trên chưa chắc.
