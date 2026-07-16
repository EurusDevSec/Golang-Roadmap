# Khu vực thực hành

Tạo mỗi bài trong một thư mục riêng:

```text
Go/practice/
├── 00-go-basics/
│   └── hello-function/main.go
├── 01-array/
│   ├── basics/main.go
│   └── statistics/main.go
└── 02-slice/
    ├── create/main.go
    ├── basics/main.go
    └── capacity/main.go
```

Mỗi thư mục con là một chương trình độc lập và chỉ có một hàm `main()`. Không đặt file `.go` trong `Go/lessons`; thư mục đó chỉ chứa tài liệu.

Definition of Done cho mỗi bài:

- `go fmt ./...` không tạo thêm thay đổi.
- `go vet ./...` không báo lỗi.
- `go test ./...` thành công.
- Lỗi được trả về hoặc log đúng chỗ, không bị bỏ qua.
- Không chứa secret trong source code.
- Có ví dụ input/output trong README nếu là project.

Không đặt lời giải trong `Go/lessons`. Mỗi chủ đề đã có `EXERCISES.md` chứa đề, input, output và trường hợp biên.
