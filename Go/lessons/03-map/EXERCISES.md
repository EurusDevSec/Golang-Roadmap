# Bài tập cơ bản 03 — Map

Chỉ dùng kiến thức: tạo map, thêm, đọc, sửa, xóa, `range` và comma-ok.

## Bài 1 — Tạo map

```go
ports := map[string]int{
	"http":  80,
	"https": 443,
}
```

Output:

```text
http=80
https=443
```

Truy cập trực tiếp bằng key để output không phụ thuộc thứ tự map.

## Bài 2 — Thêm phần tử

Thêm `ssh=22`.

```text
ssh=22
length=3
```

## Bài 3 — Cập nhật phần tử

Cho `replicas := map[string]int{"api": 2, "worker": 1}`. Đổi `api` thành `3`.

```text
before=2
after=3
```

## Bài 4 — Kiểm tra key

Kiểm tra `db` có trong map sau không:

```go
services := map[string]string{"api": "running", "db": "stopped"}
```

```text
value=stopped exists=true
```

Thử key `cache`:

```text
value= exists=false
```

Gợi ý: `value, exists := services[key]`.

## Bài 5 — Xóa key

Xóa `db` khỏi map trên.

```text
before=2
after=1
```

Gợi ý: `delete(services, "db")`.

## Bài 6 — Đếm từ

Cho `[]string{"go", "linux", "go", "docker", "go"}`.

```text
go=3
linux=1
docker=1
```

Gợi ý: tạo `counts := make(map[string]int)` rồi dùng `counts[word]++`.

## Bài 7 — Duyệt map

In tất cả key-value trong map. Chạy chương trình vài lần và quan sát thứ tự có thể thay đổi. Viết một câu kết luận về thứ tự của map.

## Hoàn thành khi

Bạn tự thêm, đọc, sửa, xóa key và làm được bài đếm từ.
