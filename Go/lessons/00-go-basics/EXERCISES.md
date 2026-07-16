# Bài tập cơ bản 00 — Làm quen lại với Go

Mỗi bài tạo một `main.go` riêng. Chỉ cần dùng `package main`, `fmt`, biến, `if` và `for`.

## Bài 1 — In thông tin

Viết chương trình in đúng ba dòng:

```text
Name: An
Role: DevOps
Language: Go
```

Gợi ý: dùng `fmt.Println` ba lần.

## Bài 2 — Khai báo biến

Tạo các biến:

```go
name := "server-01"
cpu := 4
active := true
```

Output:

```text
server=server-01 cpu=4 active=true
```

Gợi ý: dùng `fmt.Printf` với `%s`, `%d`, `%t`.

## Bài 3 — Cộng hai số

Cho:

```go
a := 10
b := 20
```

Output:

```text
sum=30
```

## Bài 4 — Kiểm tra số dương

Cho `number := -5`. Nếu lớn hơn hoặc bằng `0`, in `positive`; ngược lại in `negative`.

```text
input:  -5
output: negative
```

Thử lại với `number := 10`, output phải là `positive`.

## Bài 5 — Kiểm tra CPU

Cho `cpuUsage := 85`:

- Nếu `cpuUsage >= 80`, in `high`.
- Ngược lại, in `normal`.

```text
input:  85
output: high
```

## Bài 6 — Vòng lặp từ 1 đến 5

Output:

```text
1
2
3
4
5
```

Gợi ý: `for i := 1; i <= 5; i++`.

## Bài 7 — In số chẵn

In các số chẵn từ 1 đến 10:

```text
2 4 6 8 10
```

Gợi ý: số chẵn thỏa `i%2 == 0`.

## Bài 8 — Viết function cộng

Hoàn thành:

```go
func add(a int, b int) int {
	// TODO
}
```

Kiểm tra:

```go
fmt.Println(add(3, 5))
```

Output: `8`.

## Hoàn thành khi

Bạn tự làm được bài 1–6 và giải thích được biến, `if`, `for`, function.

