# 12 — Logging, metrics và health

## Học trọng tâm

- Structured logging với `log/slog`.
- Log level, correlation/request ID và error context.
- Metrics: counter, gauge, histogram; tránh label cardinality cao.
- `/healthz` cho process sống; `/readyz` cho khả năng nhận traffic.

## Best practices

- Không log secret, token, full config hoặc dữ liệu nhạy cảm.
- Log ở boundary, tránh cùng một lỗi bị log nhiều tầng.
- Dùng field ổn định như `service`, `operation`, `target`, `duration`.
- Không dùng URL đầy đủ/user ID ngẫu nhiên làm metric label.
- Metrics dùng để tổng hợp; logs dùng điều tra chi tiết.
- Health check phải nhanh, bounded và không tạo tải lớn lên dependency.

## Bài tập

1. Đổi log text của checker sang JSON `slog`.
2. Thêm số request thành công/thất bại và histogram latency.
3. Thêm request ID middleware.
4. Tạo `/healthz`, `/readyz` với semantics khác nhau.

