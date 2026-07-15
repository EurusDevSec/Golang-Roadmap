# 09 — Context, timeout và retry

## Học trọng tâm

- Context mang cancellation, deadline và request-scoped values.
- Đặt `context.Context` làm tham số đầu tiên; không lưu trong struct lâu dài.
- Caller quyết định timeout tổng; từng operation có thể có timeout nhỏ hơn.
- Retry cần backoff, jitter, giới hạn số lần và tôn trọng context.

## Best practices

- Luôn gọi `cancel()` sau `context.WithTimeout/Cancel`.
- Không dùng context để truyền config hoặc dependency.
- Kiểm tra `ctx.Err()` trong loop dài/chờ retry.
- Chỉ retry lỗi tạm thời: timeout, rate limit, một số 5xx.
- Không retry vô hạn; ghi nhận attempt và nguyên nhân cuối.
- Với thao tác tạo tài nguyên, dùng idempotency key hoặc kiểm tra trạng thái trước retry.

## Bài tập

1. Viết `sleep(ctx, duration) error` hủy được.
2. Gọi endpoint với deadline 2 giây.
3. Viết retry tối đa 3 lần với exponential backoff.
4. Dừng retry ngay khi context bị hủy hoặc gặp lỗi không retryable.

Checkpoint: bổ sung timeout và retry an toàn cho endpoint checker.

