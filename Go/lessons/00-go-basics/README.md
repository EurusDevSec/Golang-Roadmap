# 00 — Go basics cho automation

## Học trọng tâm

- Biến, constant, kiểu cơ bản và zero value.
- `if`, `switch`, `for`, `range`.
- Function, multiple return values, variadic parameters.
- Pointer để chia sẻ và cập nhật giá trị; Go không có pointer arithmetic.
- Scope, naming và exported identifier.

## Best practices

- Dùng `:=` trong scope nhỏ; dùng `var` khi muốn thể hiện zero value.
- Giữ function ngắn, một trách nhiệm và trả lỗi rõ ràng.
- Ưu tiên early return để giảm lồng `if`.
- Không dùng pointer chỉ để “tối ưu”; dùng khi cần mutation, identity hoặc phân biệt nil.
- Tên package ngắn, chữ thường; tên biến nhỏ theo scope, rõ theo ý nghĩa.

## Bài tập

1. Viết `bytesToMiB(bytes int64) float64`.
2. Viết `classifyStatus(code int) string` bằng `switch`.
3. Viết `parsePort(text string) (int, error)`, chỉ chấp nhận 1–65535.
4. Viết `applyPrefix(name string, prefix *string) string`; nil nghĩa là không thêm.
5. Viết chương trình nhận uptime theo giây và in dạng ngày/giờ/phút/giây.

## Checklist

- Giải thích được zero value và scope.
- Biết khi nào nên và không nên dùng pointer.
- Function validate dữ liệu trả lỗi thay vì tự kết thúc chương trình.

