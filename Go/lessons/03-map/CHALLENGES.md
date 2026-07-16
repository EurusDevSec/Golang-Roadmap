# Bài nâng cao 03 — Map

Tạo code tại `Go/practice/03-map/`.

## Bài 1 — Đếm trạng thái

Viết `func countStatus(values []string) map[string]int`.

```text
input:  [running failed running pending failed running]
output: failed=2 pending=1 running=3
```

Khi in phải sắp xếp key để output ổn định.

## Bài 2 — Gom node theo region

```go
type Node struct { Name, Region string }
func groupByRegion(nodes []Node) map[string][]string
```

```text
input:  [{api-1 sg} {db-1 vn} {api-2 sg}]
output: sg=[api-1 api-2] vn=[db-1]
```

## Bài 3 — Phần tử duy nhất đầu tiên

Viết `func firstUnique(values []int) (int, bool)`.

| Input | Output |
|---|---|
| `[4 5 4 6 5]` | `6 true` |
| `[1 1 2 2]` | `0 false` |

## Bài 4 — Gộp inventory

Viết `func merge(a, b map[string]int) map[string]int`; key trùng thì cộng giá trị và không sửa input.

```text
a:      api=2 db=1
b:      api=1 cache=3
output: api=3 cache=3 db=1
```
