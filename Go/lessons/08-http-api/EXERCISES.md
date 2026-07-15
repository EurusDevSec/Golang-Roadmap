# Bài tập 08 — HTTP API

Tạo code tại `Go/practice/08-http-api/`.

## Bài 1 — Kiểm tra endpoint

Viết `Check(ctx, client, url) Result`, trong đó `Result` có URL, StatusCode, Duration và Error.

```text
server response: HTTP 200
output: url=http://127.0.0.1:... status=200 healthy=true
```

HTTP 503 là response hợp lệ nhưng `healthy=false`; lỗi DNS/timeout nằm ở `Error`.

## Bài 2 — Giới hạn body

API trả body lớn hơn 1 MiB. Client phải dừng đọc và trả:

```text
error: response body exceeds 1048576 bytes
```

## Bài 3 — Typed API client

Endpoint `POST /jobs`, input:

```json
{"name":"backup","target":"db-01"}
```

Response `201`:

```json
{"id":"job-123","status":"queued"}
```

Tạo struct request/response; không dùng `map[string]any`. Test bằng `httptest.Server` rằng header `Authorization` được gửi nhưng không được log.

## Bài 4 — Health server

Tạo `/healthz` trả `200 {"status":"ok"}` và `/readyz` trả `503 {"status":"not_ready"}` khi dependency chưa sẵn sàng. Content-Type phải là `application/json`.
