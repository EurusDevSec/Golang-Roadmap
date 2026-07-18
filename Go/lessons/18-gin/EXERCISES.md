# Bài tập 18 — Gin thực chiến tối giản

Chỉ dùng kiến thức vừa đủ: HTTP, JSON, struct, middleware, env, health check và Docker basic.

## Bài 1 — Router đầu tiên

Tạo server Gin có 2 route:

- `GET /ping` trả `{"message":"pong"}`.
- `GET /healthz` trả `{"status":"ok"}`.

Output mong muốn:

```text
GET /ping    -> 200
GET /healthz -> 200
```

## Bài 2 — JSON typed request

Tạo `POST /jobs` nhận JSON:

```json
{ "name": "backup", "target": "db-01" }
```

Yêu cầu:

- Dùng struct cho request/response.
- Không dùng `map[string]any`.
- Trả `201` khi tạo thành công.

Output mong muốn:

```text
status=201
id=job-123
```

## Bài 3 — Validate input và lỗi

Nếu `name` rỗng hoặc thiếu `target`, trả lỗi `400`.

Yêu cầu:

- Tạo response lỗi thống nhất.
- Không để panic.

Output mong muốn:

```text
status=400
error=invalid input
```

## Bài 4 — Middleware log và recover

Tạo middleware để:

- log method, path, status, latency.
- recover khi handler panic.

Yêu cầu:

- Không log token hoặc secret.
- Mỗi request có log 1 dòng là đủ.

Output mong muốn:

```text
GET /ping status=200 latency=...
```

## Bài 5 — Config và readiness

Đọc `PORT` và `READY` từ env.

Yêu cầu:

- `GET /healthz` luôn trả `200`.
- `GET /readyz` trả `200` khi `READY=true`.
- `GET /readyz` trả `503` khi `READY` chưa sẵn sàng.

Output mong muốn:

```text
/healthz -> 200
/readyz  -> 503 or 200
```

## Bài 6 — CRUD nhỏ

Tạo API in-memory cho một tài nguyên đơn giản như `jobs` hoặc `services`.

Yêu cầu tối thiểu:

- `POST` tạo mới.
- `GET` lấy danh sách.
- `DELETE` xóa theo id.

Output mong muốn:

```text
create -> 201
list   -> 200
delete -> 200
```

## Bài 7 — Docker-ready

Chuẩn bị app để chạy trong container.

Yêu cầu:

- Port lấy từ env.
- Có `healthz`.
- Không hardcode secret.
- Chạy được bằng `go run` trước, rồi mới đóng gói Docker.

Checklist:

- [ ] Có route cơ bản
- [ ] Có JSON request/response
- [ ] Có middleware
- [ ] Có config env
- [ ] Có health check
- [ ] Có thể container hóa sau

## Gợi ý học

- Làm theo thứ tự từ Bài 1 đến Bài 7.
- Chỉ thêm thứ mới khi bài trước đã chạy được.
- Nếu vướng, quay về đọc [README.md](README.md) và làm lại từng bước.
