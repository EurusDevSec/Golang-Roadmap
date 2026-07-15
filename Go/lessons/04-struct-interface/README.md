# 04 — Struct, method và interface

## Học trọng tâm

```go
type Endpoint struct {
    Name string
    URL  string
}

func (e Endpoint) Label() string { return e.Name + ": " + e.URL }
```

- Struct mô hình hóa config, resource và kết quả kiểm tra.
- Value receiver cho kiểu nhỏ, không mutation; pointer receiver khi cần sửa hoặc tránh copy lớn.
- Embedding để tái sử dụng hành vi có chủ đích, không dùng như inheritance.
- Interface mô tả hành vi ở nơi sử dụng.

```go
type Checker interface {
    Check(ctx context.Context, target string) error
}
```

## Best practices

- “Accept interfaces, return structs” khi thật sự cần thay implementation/test double.
- Giữ interface nhỏ, thường 1–3 method.
- Không tạo interface trước khi có ít nhất hai cách dùng hợp lý.
- Constructor validate invariant và trả `(*T, error)` nếu có thể thất bại.
- Tránh struct “God object” chứa mọi dependency.

## Bài tập

1. Tạo `Server{Name, Host, Port}` và method `Address()`.
2. Tạo `Resource` interface có `ID()` và `Kind()`; implement cho VM và Bucket.
3. Tạo `Store` interface nhỏ để lưu/đọc server, kèm in-memory implementation.
4. Viết constructor từ chối host rỗng hoặc port sai.

Checkpoint: chương trình inventory nhóm resource theo `Kind` và in summary.

