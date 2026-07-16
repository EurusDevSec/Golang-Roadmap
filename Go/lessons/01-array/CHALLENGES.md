# Bài nâng cao 01 — Array

Tạo code tại `Go/practice/01-array/`. Mỗi bài dùng array đúng kích thước, không đổi sang slice trừ khi đề yêu cầu.

## Bài 1 — Thống kê CPU 5 phút

Cho `[5]float64{20.5, 35, 80.5, 70, 44}`. In min, max và average với 2 chữ số thập phân.

```text
min=20.50 max=80.50 avg=50.00
```

## Bài 2 — Đảo thứ tự node

Viết `func reverse(nodes *[5]string)` sửa tại chỗ.

```text
input:  [node-1 node-2 node-3 node-4 node-5]
output: [node-5 node-4 node-3 node-2 node-1]
```

Không tạo array phụ.

## Bài 3 — Tìm lần đầu vượt ngưỡng

Viết `func firstAbove(values [6]int, threshold int) (index int, ok bool)`.

| Values | Threshold | Output |
|---|---:|---|
| `[40 55 72 91 60 30]` | `80` | `3 true` |
| `[40 55 72 91 60 30]` | `100` | `-1 false` |

## Bài 4 — Xoay lịch trực

Xoay phải `[7]string` một vị trí bằng đúng một biến tạm.

```text
input:  [Mon Tue Wed Thu Fri Sat Sun]
output: [Sun Mon Tue Wed Thu Fri Sat]
```
