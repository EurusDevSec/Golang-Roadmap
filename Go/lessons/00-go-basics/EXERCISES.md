# Bài tập 00 — Go basics

Tạo code tại `Go/practice/00-go-basics/`.

## Bài 1 — Đổi byte sang MiB

Viết `func bytesToMiB(value int64) (float64, error)`. Giá trị âm trả lỗi.

| Input | Output |
|---:|---:|
| `1048576` | `1.00` |
| `1572864` | `1.50` |
| `0` | `0.00` |
| `-1` | error |

## Bài 2 — Phân loại HTTP status

Viết `func classifyStatus(code int) string`.

| Input | Output |
|---:|---|
| `200` | `success` |
| `302` | `redirect` |
| `404` | `client_error` |
| `503` | `server_error` |
| `99` | `invalid` |

## Bài 3 — Validate port

CLI nhận một đối số là port.

```text
> go run . 8080
valid port: 8080

> go run . 70000
error: port must be between 1 and 65535
```

Exit code: `0` nếu hợp lệ, khác `0` nếu sai hoặc thiếu input.

## Bài 4 — Format uptime

Viết `func formatUptime(seconds int64) string`.

| Input | Output |
|---:|---|
| `59` | `0d 00h 00m 59s` |
| `3661` | `0d 01h 01m 01s` |
| `90061` | `1d 01h 01m 01s` |

Không chấp nhận số âm.

