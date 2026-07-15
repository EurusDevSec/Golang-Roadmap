# Bài tập 10 — Concurrency

Tạo code tại `Go/practice/10-concurrency/` và luôn chạy `go test -race ./...`.

## Bài 1 — Worker pool

Xử lý các job `[1 2 3 4 5 6]` bằng tối đa 2 worker. Mỗi job trả bình phương.

```text
input:  [1 2 3 4 5 6]
output: [1 4 9 16 25 36]
```

Thứ tự kết quả cuối phải trùng input dù thứ tự hoàn thành khác nhau.

## Bài 2 — Giới hạn HTTP concurrency

Kiểm tra 20 URL; test server đếm số request đang chạy. Kết quả bắt buộc:

```text
total=20 max_concurrent<=4
```

## Bài 3 — Hủy batch

Có 10 job, mỗi job 500ms, 2 worker; context timeout 100ms.

```text
completed<10
error=context deadline exceeded
elapsed<300ms
```

Tất cả worker phải thoát.

## Bài 4 — Counter race

100 goroutine, mỗi goroutine tăng counter 1.000 lần.

```text
output: 100000
```

Bảo vệ bằng mutex hoặc atomic; race detector không báo lỗi.

