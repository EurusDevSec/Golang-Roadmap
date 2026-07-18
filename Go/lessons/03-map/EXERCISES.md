# Bài tập cơ bản 03 — Map

Chỉ dùng kiến thức: tạo map, thêm, đọc, sửa, xóa, `range` và comma-ok.

## Quy ước output (quan trọng)

- In theo dạng `key=value`, mỗi dòng một cặp.
- Không dùng `fmt.Println(map)` để nộp kết quả, vì sẽ ra dạng `map[...]` khó so sánh.
- Với bài có ghi "thứ tự có thể khác": chỉ cần đúng dữ liệu, không bắt buộc đúng thứ tự dòng.
- Nếu muốn output ổn định để tự check dễ hơn: in theo key cố định (ví dụ `http`, `https`) thay vì `range` trực tiếp.

## Bài 1 — Tạo map

```go
ports := map[string]int{
	"http":  80,
	"https": 443,
}
```

Output mẫu (một cách đúng):

```text
http=80
https=443
```

Tiêu chí đạt: in đúng 2 dòng trên, đúng format `key=value`.

## Bài 2 — Thêm phần tử

Thêm `ssh=22`.

Output mẫu:

```text
ssh=22
length=3
```

Tiêu chí đạt: có `ssh=22` và độ dài map sau khi thêm là `3`.

## Bài 3 — Cập nhật phần tử

Cho `replicas := map[string]int{"api": 2, "worker": 1}`. Đổi `api` thành `3`.

Output mẫu:

```text
before=2
after=3
```

Tiêu chí đạt: giá trị `api` đổi từ `2` sang `3`.

## Bài 4 — Kiểm tra key

Kiểm tra `db` có trong map sau không:

```go
services := map[string]string{"api": "running", "db": "stopped"}
```

Output mẫu (key `db`):

```text
value=stopped exists=true
```

Thử key `cache`:

Output mẫu (key `cache`):

```text
value= exists=false
```

Tiêu chí đạt: phân biệt rõ key có tồn tại và không tồn tại bằng `exists`.
Gợi ý: `value, exists := services[key]`.

## Bài 5 — Xóa key

Xóa `db` khỏi map trên.

Output mẫu:

```text
before=2
after=1
```

Gợi ý: `delete(services, "db")`.

## Bài 6 — Đếm từ

Cho `[]string{"go", "linux", "go", "docker", "go"}`.

Output mẫu (thứ tự dòng có thể khác):

```text
go=3
linux=1
docker=1
```

Tiêu chí đạt: đúng số đếm từng từ, không cần đúng thứ tự dòng.

Gợi ý: tạo `counts := make(map[string]int)` rồi dùng `counts[word]++`.

## Bài 7 — Duyệt map

In tất cả key-value trong map. Chạy chương trình vài lần và quan sát thứ tự có thể thay đổi.

Mẫu kết luận:

```text
Thu tu duyet map trong Go khong co tinh on dinh, co the thay doi giua cac lan chay.
```

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

Output mẫu:

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

Output mẫu:

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

Output mẫu (thứ tự có thể khác):

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

Output mẫu:

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

Output mẫu:

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

Output mẫu:

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

Output mẫu:

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

Output mẫu:

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

Output mẫu:

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

## Cách tự chấm nhanh

- Đúng format: mỗi dòng là `ten_truong=gia_tri`.
- Đúng dữ liệu: số/count/trạng thái đúng với đề.
- Với bài có map `range`: chấp nhận khác thứ tự dòng.
- Nếu bạn in trực tiếp map ra `map[...]` thì coi như chưa đạt phần format output.
