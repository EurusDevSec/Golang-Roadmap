# Khu vực thực hành

Tạo mỗi bài trong một thư mục riêng:

```text
Go/practice/
├── endpoint-checker/
│   ├── main.go
│   └── main_test.go
└── config-validator/
    ├── main.go
    └── main_test.go
```

Definition of Done cho mỗi bài:

- `go fmt ./...` không tạo thêm thay đổi.
- `go vet ./...` không báo lỗi.
- `go test ./...` thành công.
- Lỗi được trả về hoặc log đúng chỗ, không bị bỏ qua.
- Không chứa secret trong source code.
- Có ví dụ input/output trong README nếu là project.

Không đặt lời giải trong `Go/lessons`. Mỗi chủ đề đã có `EXERCISES.md` chứa đề, input, output và trường hợp biên.
