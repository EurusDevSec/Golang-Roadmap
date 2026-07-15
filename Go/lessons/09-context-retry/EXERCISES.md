# Bài tập 09 — Context và retry

Tạo code tại `Go/practice/09-context-retry/`.

## Bài 1 — Sleep hủy được

Viết `func Sleep(ctx context.Context, d time.Duration) error`.

| Context | Duration | Output |
|---|---:|---|
| còn hiệu lực | `10ms` | `nil` sau khoảng 10ms |
| cancel ngay | `10s` | `context.Canceled` ngay |
| deadline 20ms | `10s` | `context.DeadlineExceeded` |

Không leak timer/goroutine.

## Bài 2 — Retry policy

Operation lỗi tạm thời 2 lần rồi thành công:

```text
attempt=1 error=temporary
attempt=2 error=temporary
attempt=3 success
```

Operation lỗi validation phải dừng ở attempt 1. Tối đa 3 attempts, không phải 3 retries.

## Bài 3 — Context dừng backoff

Backoff dự kiến `100ms, 200ms, 400ms`; context bị hủy sau `150ms`.

```text
attempts=2
error=context canceled
elapsed<250ms
```

## Bài 4 — Idempotent retry

Chỉ retry `GET`, `HEAD` và `PUT`. Với `POST` chỉ retry khi có `Idempotency-Key`.

| Method | Header | Retry? |
|---|---|---|
| GET | none | yes |
| POST | none | no |
| POST | Idempotency-Key=abc | yes |

