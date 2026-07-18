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

---

## Bài tập mở rộng — Map cho DevOps/Cloud/Backend

Chỉ dùng kiến thức: array, slice, map, `range`, comma-ok, `delete`.
Không cần thuật toán phức tạp.

## Bài 8 — Port mặc định theo service

Cho map:

```go
defaultPorts := map[string]int{
	"http":   80,
	"https":  443,
	"ssh":    22,
	"postgres": 5432,
}
```

Yêu cầu:

- In ra `https=443`.
- Kiểm tra key `redis` có tồn tại hay không bằng comma-ok.

Output mong muốn:

```text
https=443
redis_exists=false
```

## Bài 9 — Trạng thái service

Cho map:

```go
services := map[string]string{
	"api": "running",
	"db":  "stopped",
	"mq":  "running",
}
```

Yêu cầu:

- Đổi `db` thành `running`.
- In trước và sau khi đổi.

Output mong muốn:

```text
before=stopped
after=running
```

## Bài 10 — Đếm env trong danh sách pod

Cho slice:

```go
envs := []string{"prod", "staging", "prod", "dev", "prod", "staging"}
```

Yêu cầu:

- Đếm số lần xuất hiện của từng env bằng map.

Output mong muốn (thứ tự có thể khác):

```text
prod=3
staging=2
dev=1
```

## Bài 11 — Whitelist IP (map bool)

Cho slice:

```go
allowIPs := []string{"10.0.0.1", "10.0.0.2", "10.0.0.1"}
checkIP := "10.0.0.2"
```

Yêu cầu:

- Dùng `map[string]bool` để tạo whitelist.
- Kiểm tra `checkIP` có được phép hay không.

Output mong muốn:

```text
allowed=true
unique_ips=2
```

## Bài 12 — Group log theo level

Cho slice:

```go
levels := []string{"INFO", "ERROR", "INFO", "WARN", "ERROR"}
```

Yêu cầu:

- Tạo `map[string][]string` để gom log message giả lập theo level.
- Bạn có thể tự gán message dạng `msg-1`, `msg-2`, ... khi duyệt slice.

Ví dụ output:

```text
INFO=[msg-1 msg-3]
ERROR=[msg-2 msg-5]
WARN=[msg-4]
```

## Bài 13 — Merge config mặc định và config môi trường

Cho 2 map:

```go
defaults := map[string]string{
	"region": "ap-southeast-1",
	"log_level": "info",
	"timeout": "30s",
}

override := map[string]string{
	"log_level": "debug",
	"timeout": "10s",
}
```

Yêu cầu:

- Tạo map mới `finalCfg`.
- Copy toàn bộ `defaults` vào `finalCfg`, sau đó ghi đè bằng `override`.
- Không sửa trực tiếp `defaults`.

Output mong muốn:

```text
region=ap-southeast-1
log_level=debug
timeout=10s
```

## Bài 14 — Xóa key nhạy cảm trước khi in

Cho map:

```go
config := map[string]string{
	"db_user": "admin",
	"db_password": "secret",
	"api_key": "abc-123",
	"region": "ap-southeast-1",
}
```

Yêu cầu:

- Xóa `db_password` và `api_key` trước khi in.
- In `before` và `after` theo số lượng phần tử.

Output mong muốn:

```text
before=4
after=2
```

## Bài 15 — HTTP status summary

Cho slice:

```go
statuses := []int{200, 200, 500, 503, 200, 404, 503}
```

Yêu cầu:

- Đếm tần suất status code bằng `map[int]int`.
- In ra số lượng của `200`, `500`, `503`.

Output mong muốn:

```text
200=3
500=1
503=2
```

## Bài 16 — Multi-tenant tài nguyên

Cho slice tenant:

```go
tenants := []string{"team-a", "team-b", "team-a", "team-c", "team-b", "team-a"}
```

Yêu cầu:

- Đếm số resource theo tenant.
- Tìm tenant có nhiều resource nhất.

Ví dụ output:

```text
team-a=3
team-b=2
team-c=1
max_tenant=team-a max_count=3
```

## Gợi ý làm nhanh

- Bài 8, 9, 13, 14: luyện đọc/sửa/xóa key rất quan trọng khi xử lý config thực tế.
- Bài 10, 15, 16: luyện map làm counter cho metric/log/monitoring.
- Bài 11: luyện map như set để lọc trùng và check membership.
- Bài 12: luyện map chứa slice, rất hay gặp khi gom dữ liệu theo nhóm.
