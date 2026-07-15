# Bài tập 11 — Testing

Tạo code tại `Go/practice/11-testing/`.

## Bài 1 — Table test cho port

Test `ParsePort(string) (int, error)` với bảng:

| Input | Value | Error? |
|---|---:|---|
| `"80"` | `80` | no |
| `"65535"` | `65535` | no |
| `"0"` | `0` | yes |
| `"abc"` | `0` | yes |
| `""` | `0` | yes |

Mỗi case dùng `t.Run`.

## Bài 2 — Config file test

Dùng `t.TempDir()` tạo:

- `valid.json`: loader thành công.
- `invalid.json`: JSON hỏng, trả lỗi.
- đường dẫn không tồn tại: `errors.Is(err, os.ErrNotExist)` là true.

Test không đọc file cố định trong máy.

## Bài 3 — HTTP client test

Dùng `httptest.Server` tạo 4 subtests:

| Case | Server | Expected |
|---|---|---|
| success | 200 JSON | decode thành công |
| server_error | 500 | typed status error |
| invalid_json | 200 body `{` | decode error |
| timeout | chậm 200ms, deadline 20ms | deadline exceeded |

## Bài 4 — Fuzz parser

Fuzz parser dòng log. Điều kiện: không panic với bất kỳ `[]byte`; dòng hợp lệ phải round-trip được.

