# Bài tập cơ bản 01 — Array

Chỉ dùng kiến thức: khai báo array, chỉ số, `len`, `for`, `range`.

## Bài 1 — Tạo và in array

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

Output:

```text
[1 2 3 4 5]
```

## Bài 2 — Lấy phần tử

Dùng array trên và in phần tử đầu, phần tử thứ ba, phần tử cuối.

```text
first=1
third=3
last=5
```

Gợi ý: chỉ số bắt đầu từ `0`; phần tử cuối là `numbers[len(numbers)-1]`.

## Bài 3 — Thay đổi phần tử

Cho:

```go
numbers := [3]int{10, 20, 30}
```

Đổi phần tử `20` thành `99`.

```text
before: [10 20 30]
after:  [10 99 30]
```

## Bài 4 — Duyệt array

Cho `[4]string{"api", "db", "cache", "worker"}`. In:

```text
0: api
1: db
2: cache
3: worker
```

Gợi ý: `for index, value := range services`.

## Bài 5 — Tính tổng

Cho `[5]int{2, 4, 6, 8, 10}`.

```text
sum=30
```

Gợi ý: tạo `sum := 0`, sau đó cộng từng phần tử.

## Bài 6 — Đếm số lớn hơn 5

Cho `[6]int{2, 8, 1, 9, 5, 7}`.

```text
count=3
```

## Bài 7 — Tìm số lớn nhất

Cho `[5]int{4, 10, 3, 8, 6}`.

```text
max=10
```

Gợi ý: bắt đầu bằng `max := numbers[0]`, sau đó so sánh các phần tử.

## Hoàn thành khi

Bạn hiểu `[5]int` có đúng 5 phần tử và tự làm được bài 1–6.

