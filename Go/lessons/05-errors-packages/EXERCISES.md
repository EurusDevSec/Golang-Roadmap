# Bài tập 05 — Error và package

Tạo code tại `Go/practice/05-errors-packages/`.

## Bài 1 — Đọc token

Viết `func LoadToken(path string) (string, error)` trong package `config`.

```text
file content: "  abc123\n"
output: abc123
```

File không tồn tại: lỗi phải giữ được `os.ErrNotExist` qua `errors.Is`. File chỉ có khoảng trắng: trả `ErrEmptyToken`.

## Bài 2 — ValidationError

Tạo error chứa `Field` và `Message`.

```text
input:  address="", timeout=-1
output: field=address message=must not be empty
```

Caller phải đọc được field bằng `errors.As`.

## Bài 3 — CLI exit code

Viết `run(args []string) error`; `main` ánh xạ lỗi:

| Trường hợp | Exit code |
|---|---:|
| Thành công | `0` |
| Input/config sai | `2` |
| Lỗi runtime | `1` |

Không dùng `log.Fatal` bên trong package `config` hoặc `app`.

