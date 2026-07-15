# 08 — HTTP client và API

## Học trọng tâm

- Tái sử dụng `http.Client`; không tạo client mới cho mỗi request.
- Tạo request bằng `http.NewRequestWithContext`.
- Luôn đóng response body và kiểm tra status code.
- Server: route, validate input, status code, JSON response và middleware nhỏ.

```go
resp, err := client.Do(req)
if err != nil { return err }
defer resp.Body.Close()
```

## Best practices

- Client luôn có timeout hoặc deadline từ context.
- Giới hạn body trước khi đọc toàn bộ.
- Không retry mọi lỗi; chỉ retry lỗi tạm thời và request idempotent.
- Không log Authorization header, cookie hoặc response chứa secret.
- Dùng typed request/response thay cho `map[string]any` khi schema đã biết.
- Server đặt read/write/idle timeout và shutdown có deadline.

## Bài tập

1. Health checker trả URL, status, latency và error.
2. API client có base URL và token được inject qua constructor.
3. HTTP server có `/healthz` và `/readyz`.
4. Mock server bằng `httptest.Server` để test status 200, 500 và timeout.

Checkpoint: CLI kiểm tra danh sách endpoint và xuất JSON hợp lệ.

