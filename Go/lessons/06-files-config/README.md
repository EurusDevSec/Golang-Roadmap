# 06 — File, JSON, YAML và config

## Học trọng tâm

- `os.ReadFile` cho file nhỏ; `bufio.Scanner`/`Reader` cho stream lớn.
- `io.Reader`/`io.Writer` giúp code dễ test và dùng lại.
- JSON tags, encode/decode và validate sau khi decode.
- YAML thường dùng thư viện `gopkg.in/yaml.v3`; khóa version trong `go.mod`.
- Thứ tự config nên rõ: default → file → environment → flag.

```go
type Config struct {
    Address string        `json:"address" yaml:"address"`
    Timeout time.Duration `json:"-" yaml:"-"`
}
```

## Best practices

- Không lưu token/password trong repo hoặc log.
- Giới hạn kích thước input nếu đọc dữ liệu không tin cậy.
- Decode nghiêm ngặt khi typo trong config có thể gây sự cố.
- Validate toàn bộ config trước khi khởi động công việc.
- Ghi file quan trọng theo kiểu atomic: file tạm → sync/close → rename.
- Truyền `io.Reader` vào parser thay vì để parser tự mở file.

## Bài tập

1. Đọc danh sách endpoint từ JSON và validate URL.
2. Đọc file log theo dòng, đếm `ERROR`, `WARN`, `INFO`.
3. Viết config loader hỗ trợ default, environment và file.
4. Chuyển inventory JSON thành CSV qua `io.Reader`/`io.Writer`.

Checkpoint: config validator trả exit code khác 0 và chỉ rõ field sai.

