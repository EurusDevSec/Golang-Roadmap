# Bài tập 02 — Slice

Tạo code tại `Go/practice/02-slice/`.

## Bài 1 — Lọc port hợp lệ

Viết `func validPorts(values []int) []int`. Không sửa input.

```text
input:  [80 -1 443 0 65535 65536 8080]
output: [80 443 65535 8080]
```

## Bài 2 — Xóa endpoint theo index

Viết `func removeAt(values []string, index int) ([]string, error)`, giữ thứ tự.

```text
input:  [api db cache], index=1
output: [api cache]
```

Index `-1`, `3` hoặc slice rỗng phải trả lỗi, không panic.

## Bài 3 — Chia batch

Viết `func chunk(values []string, size int) ([][]string, error)`.

```text
input:  [a b c d e], size=2
output: [[a b] [c d] [e]]
```

`size <= 0` trả lỗi. Input rỗng trả slice rỗng.

## Bài 4 — Bản sao độc lập

Viết `func clone(values []string) []string`. Với input `[dev staging prod]`, sửa `clone[0] = "local"` không được làm input thay đổi.

```text
original: [dev staging prod]
clone:    [local staging prod]
```

