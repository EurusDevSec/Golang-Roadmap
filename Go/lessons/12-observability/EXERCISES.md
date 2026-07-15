# Bài tập 12 — Observability

Tạo code tại `Go/practice/12-observability/`.

## Bài 1 — Structured log

Với một health check thất bại, ghi đúng một JSON log:

```json
{"level":"ERROR","msg":"health check failed","operation":"check","target":"api","status":503}
```

Thời gian có thể khác. Không có field token, authorization hoặc response body.

## Bài 2 — Metrics

Sau 3 request thành công và 2 request lỗi, metrics phải biểu diễn:

```text
requests_total{result="success"} 3
requests_total{result="error"} 2
```

Không dùng URL đầy đủ làm label. Thêm histogram latency với đơn vị giây.

## Bài 3 — Health và readiness

| Trạng thái | `/healthz` | `/readyz` |
|---|---:|---:|
| process sống, DB sẵn sàng | 200 | 200 |
| process sống, DB mất kết nối | 200 | 503 |

Readiness check timeout tối đa 500ms.

## Bài 4 — Request ID

Nếu request có `X-Request-ID: req-123`, response và log dùng `req-123`. Nếu thiếu, middleware sinh ID không rỗng và trả lại trong response header.

