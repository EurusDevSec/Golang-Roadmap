# 13 — Service runtime và graceful shutdown

## Học trọng tâm

- Dependency wiring trong `main`.
- Lifecycle: load config → init dependency → start → signal → drain → close.
- HTTP server timeouts và `Shutdown(ctx)`.
- Background worker nhận context và báo lỗi về supervisor.

## Best practices

- Startup thất bại nhanh nếu config/dependency bắt buộc không hợp lệ.
- Shutdown có deadline; không chờ vô hạn.
- Ngừng nhận việc mới trước khi chờ việc đang chạy.
- Cleanup theo thứ tự ngược với startup.
- Không gọi `os.Exit` hoặc `log.Fatal` trong package thư viện vì bỏ qua defer.
- Không chạy background goroutine mà không theo dõi lỗi/lifecycle.

## Bài tập

1. HTTP server shutdown trong tối đa 10 giây.
2. Worker dừng nhận job khi có SIGTERM nhưng hoàn tất job hiện tại.
3. Trả lỗi startup nếu port đang được sử dụng.
4. Viết test xác nhận shutdown không leak goroutine.

