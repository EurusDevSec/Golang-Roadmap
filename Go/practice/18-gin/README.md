# Khu vực thực hành — Gin tối giản

Mỗi bài là một app nhỏ độc lập. Mục tiêu là học nhanh phần dùng thật, không nhảy sang kiến trúc phức tạp.

```text
Go/practice/18-gin/
├── 00-router/
│   └── main.go
├── 01-json-binding/
│   └── main.go
├── 02-middleware/
│   └── main.go
├── 03-config-health/
│   └── main.go
└── 04-crud-memory/
    └── main.go
```

## Cách học

1. Chạy từng file bằng `go run`.
2. Đọc route, request, response và middleware.
3. Sửa handler để thử status code khác.
4. Tự thêm log hoặc field JSON để quan sát output.

## Thứ tự nên làm

- `00-router`: router, method, JSON response.
- `01-json-binding`: typed request/response.
- `02-middleware`: log và recover.
- `03-config-health`: env, healthz, readyz.
- `04-crud-memory`: CRUD nhỏ bằng map.
