# 07 — CLI, environment, process và signal

## Học trọng tâm

- `flag` đủ cho CLI nhỏ; chỉ thêm framework khi cần subcommand/UX phức tạp.
- `os.Args`, environment variables và exit code.
- `os/exec` để chạy process; truyền argument riêng, không ghép shell command từ input.
- Signal và graceful cancellation với `signal.NotifyContext`.

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()
```

## Best practices

- `main` gọi `run() error`, log lỗi rồi chọn exit code tại một nơi.
- Ghi dữ liệu bình thường ra stdout, diagnostics ra stderr.
- Exit code `0` là thành công; mã khác 0 phải có ý nghĩa ổn định.
- Không truyền secret qua command line nếu có lựa chọn file/stdin/env an toàn hơn.
- Dùng `exec.CommandContext` để process con bị hủy theo timeout.
- Không dùng `sh -c` với input người dùng.

## Bài tập

1. CLI `envcheck` kiểm tra các biến bắt buộc.
2. CLI `disk-report` nhận format `text|json`.
3. Chạy một command với timeout 3 giây và thu stdout/stderr riêng.
4. Xử lý Ctrl+C để dừng vòng lặp công việc sạch sẽ.

