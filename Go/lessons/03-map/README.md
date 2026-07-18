# Map trong Go

## Mục tiêu

- Khởi tạo, đọc, ghi, xóa và duyệt map.
- Phân biệt giá trị không tồn tại với zero value.
- Biết cách dùng map làm set và bộ đếm.
- Tránh phụ thuộc vào thứ tự duyệt map.

## 1. Kiến thức cốt lõi

Map lưu cặp key-value với kiểu `map[K]V`. Key phải là kiểu có thể so sánh bằng `==`.

```go
ages := map[string]int{
    "An":   20,
    "Binh": 22,
}

scores := make(map[string]float64)
scores["Go"] = 9.5
```

Zero value của map là `nil`:

```go
var users map[string]int
fmt.Println(users["missing"]) // 0: đọc được
// users["An"] = 1           // panic: không thể ghi vào nil map
```

Dùng `make` hoặc map literal trước khi ghi.

## 2. Kiểm tra key tồn tại

Đọc một key không tồn tại trả về zero value. Dùng comma-ok để phân biệt:

```go
age, ok := ages["An"]
if !ok {
    fmt.Println("không tìm thấy")
} else {
    fmt.Println(age)
}
```

Khi chỉ cần kiểm tra tồn tại:

```go
if _, ok := ages["An"]; ok {
    fmt.Println("đã tồn tại")
}
```

## 3. Xóa, làm rỗng và duyệt

```go
delete(ages, "An") // an toàn ngay cả khi key không tồn tại
clear(ages)        // Go 1.21+: xóa tất cả entries

for name, age := range ages {
    fmt.Println(name, age)
}
```

Thứ tự duyệt map không được đảm bảo. Nếu cần kết quả ổn định, lấy key ra slice rồi sắp xếp:

```go
keys := make([]string, 0, len(ages))
for key := range ages {
    keys = append(keys, key)
}
slices.Sort(keys)

for _, key := range keys {
    fmt.Println(key, ages[key])
}
```

## 4. Mẫu sử dụng phổ biến

### Bộ đếm

Không cần kiểm tra key vì zero value của `int` là `0`:

```go
counts := make(map[string]int)
for _, word := range words {
    counts[word]++
}
```

### Set

```go
seen := make(map[string]bool)
seen["go"] = true

if seen["go"] {
    fmt.Println("đã gặp")
}
```

`map[string]bool` dễ đọc hơn cho người mới học: `true` là có, `false` là không có.

### Map chứa slice

```go
groups := make(map[string][]string)
groups["backend"] = append(groups["backend"], "An")
```

Không cần khởi tạo slice riêng vì có thể `append` vào nil slice.

## 5. Map và hàm

Khi truyền map vào hàm, việc thêm, sửa hoặc xóa entry có thể được phía gọi nhìn thấy:

```go
func addUser(users map[string]int, name string, age int) {
    users[name] = age
}
```

Nếu không muốn thay đổi input, tạo bản sao:

```go
func clone[K comparable, V any](source map[K]V) map[K]V {
    result := make(map[K]V, len(source))
    for key, value := range source {
        result[key] = value
    }
    return result
}
```

Đây là shallow copy. Nếu value chứa slice, map hoặc pointer, dữ liệu lồng nhau vẫn có thể được chia sẻ.

## 6. Best practices

1. Khởi tạo map bằng `make` hoặc literal trước khi ghi.
2. Dùng comma-ok khi zero value và trạng thái “không tồn tại” có ý nghĩa khác nhau.
3. Không phụ thuộc vào thứ tự của `range` trên map; sắp xếp key khi cần output ổn định.
4. Với mức cơ bản, dùng `map[T]bool` để biểu diễn set cho dễ hiểu.
5. Preallocate bằng `make(map[K]V, expectedSize)` khi biết gần đúng số entry.
6. Không giữ reference tới map nội bộ nếu API cần bảo vệ dữ liệu; trả về bản sao.
7. Nhớ rằng clone thông thường chỉ là shallow copy.
8. Map không an toàn cho nhiều goroutine cùng ghi; phải đồng bộ bằng mutex hoặc thiết kế ownership rõ ràng.
9. Đặt tên theo ý nghĩa như `usersByID`, `countByWord`, thay vì `myMap`.

## 7. Bài tập

### Mức dễ

#### Bài 1: Danh bạ

Tạo `map[string]string` lưu tên và số điện thoại. Thêm 3 người, cập nhật 1 số, xóa 1 người và in số entry còn lại.

#### Bài 2: Tra cứu an toàn

Viết `findPrice(prices map[string]float64, product string) (float64, bool)` và kiểm tra cả sản phẩm có giá `0`.

#### Bài 3: Đếm từ

Viết `countWords(words []string) map[string]int`.

Ví dụ: `["go", "java", "go"]` thành `map[go:2 java:1]`.

#### Bài 4: Set phần tử duy nhất

Viết `toSet(values []int) map[int]bool` và kiểm tra một số có thuộc set không.

### Mức dễ đến vừa

#### Bài 5: Đảo key-value

Viết `invert(source map[string]int) (map[int]string, bool)`.

Trả `false` nếu hai key có cùng value vì khi đảo sẽ mất dữ liệu.

#### Bài 6: Gom nhóm theo độ dài

Viết `groupByLength(words []string) map[int][]string`.

Ví dụ: `["go", "map", "hi"]` thành `map[2:[go hi] 3:[map]]`.

#### Bài 7: Gộp bộ đếm

Viết `mergeCounts(a, b map[string]int) map[string]int`. Kết quả phải là map mới và không sửa `a`, `b`.

#### Bài 8: Phần tử xuất hiện đầu tiên một lần

Viết `firstUnique(values []int) (int, bool)` trả phần tử đầu tiên có số lần xuất hiện bằng `1`.

Ví dụ: `[4, 5, 4, 6, 5]` trả `(6, true)`.

## 8. Gợi ý

<details>
<summary>Mở gợi ý</summary>

- Bài 2: dùng `price, ok := prices[product]`.
- Bài 3: `counts[word]++` hoạt động cả khi key chưa tồn tại.
- Bài 5: trước khi gán vào map kết quả, kiểm tra value đã trở thành key chưa.
- Bài 6: append trực tiếp vào `groups[len(word)]`.
- Bài 7: sao chép `a`, sau đó cộng từng entry của `b`.
- Bài 8: duyệt lần một để đếm, lần hai để tìm phần tử đầu tiên có count bằng `1`.

</details>

## Checklist

- Tôi biết vì sao không được ghi vào nil map.
- Tôi dùng được comma-ok.
- Tôi không phụ thuộc vào thứ tự duyệt map.
- Tôi biết dùng map làm counter và set.
- Tôi hiểu clone map là shallow copy.
- Tôi hoàn thành ít nhất 6/8 bài tập.
