# 05 — Error, package và module

## Học trọng tâm

- Error là value; trả lỗi lên caller để caller quyết định.
- Bọc lỗi bằng `%w`; kiểm tra bằng `errors.Is`/`errors.As`.
- Sentinel error cho trạng thái ổn định; custom error khi caller cần dữ liệu.
- `panic` chỉ dành cho trạng thái chương trình không thể tiếp tục, không dùng cho input/API lỗi.
- Module là đơn vị version; package là đơn vị tổ chức code.

```go
data, err := os.ReadFile(path)
if err != nil {
    return fmt.Errorf("read config %q: %w", path, err)
}
```

## Best practices

- Thêm context khi bọc lỗi nhưng không lặp từ “failed” ở mọi tầng.
- Chỉ log lỗi một lần tại boundary; tầng dưới trả lỗi.
- Không bỏ lỗi bằng `_` nếu chưa có lý do rõ ràng.
- Tránh package tên `utils`, `common`; đặt theo domain như `config`, `health`.
- Giữ `main` mỏng: parse input, wiring dependency, gọi application, chọn exit code.

## Bài tập

1. Viết `loadToken(path string) (string, error)` với lỗi có context.
2. Tạo `ErrNotFound` và dùng `errors.Is`.
3. Tạo `ValidationError{Field, Message}` và dùng `errors.As`.
4. Tách một CLI nhỏ thành `cmd`, `internal/config`, `internal/app`.

