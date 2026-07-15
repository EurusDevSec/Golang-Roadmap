# Golang Roadmap

Repo học Go theo từng chủ đề, ưu tiên kiến thức cốt lõi, best practices và bài tập thực hành.

## Cấu trúc

```text
Go/
├── lessons/          # Lý thuyết và bài tập bằng Markdown
│   ├── 01-array/
│   ├── 02-slice/
│   └── 03-map/
└── examples/         # Mã Go có thể chạy
    ├── array/
    ├── slice/
    └── map/
```

## Thứ tự học

1. [Array](Go/lessons/01-array/README.md)
2. [Slice](Go/lessons/02-slice/README.md)
3. [Map](Go/lessons/03-map/README.md)

## Chạy ví dụ

Từ thư mục gốc:

```bash
go run ./Go/examples/array
go run ./Go/examples/slice
go run ./Go/examples/map
```

Mỗi bài tập nên được làm trong một thư mục riêng để các file `package main` không xung đột nhau.

