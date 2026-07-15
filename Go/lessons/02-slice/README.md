# Slice trong Go

## Mục tiêu

- Tạo, đọc, sửa và mở rộng slice.
- Hiểu `len`, `cap` và backing array.
- Biết cách sao chép, xóa và truyền slice vào hàm.
- Tránh lỗi chia sẻ dữ liệu ngoài ý muốn.

## 1. Kiến thức cốt lõi

Slice có kiểu `[]T`, không chứa kích thước trong kiểu:

```go
numbers := []int{10, 20, 30}
empty := make([]int, 0)
buffer := make([]byte, 0, 64)
```

- `len(s)`: số phần tử đang dùng.
- `cap(s)`: số phần tử có thể chứa từ vị trí đầu slice đến cuối backing array.
- Zero value của slice là `nil`; có thể đọc, `len`, `range` và `append` an toàn.

```go
var values []int
fmt.Println(values == nil, len(values)) // true 0
values = append(values, 5)
```

### Cắt slice

```go
numbers := []int{10, 20, 30, 40, 50}
part := numbers[1:4] // [20 30 40], gồm chỉ số 1 nhưng không gồm 4
```

Hai slice được cắt từ cùng dữ liệu thường dùng chung backing array:

```go
part[0] = 99
fmt.Println(numbers) // [10 99 30 40 50]
```

### Thêm phần tử

Luôn nhận lại kết quả của `append` vì nó có thể cấp phát backing array mới:

```go
numbers = append(numbers, 40)
numbers = append(numbers, 50, 60)
```

### Sao chép độc lập

```go
source := []int{1, 2, 3}
cloned := append([]int(nil), source...)

// Hoặc:
cloned = make([]int, len(source))
copy(cloned, source)
```

### Duyệt và sửa

```go
for i, value := range numbers {
    fmt.Println(i, value)
}

for i := range numbers {
    numbers[i] *= 2
}
```

`value` là bản sao; muốn sửa phần tử phải dùng `numbers[i]`.

## 2. Truyền slice vào hàm

Slice descriptor được truyền theo giá trị nhưng vẫn trỏ đến backing array chung. Vì vậy sửa phần tử có thể ảnh hưởng dữ liệu gốc:

```go
func double(values []int) {
    for i := range values {
        values[i] *= 2
    }
}
```

Nếu hàm `append` và cần thay đổi độ dài ở phía gọi, hãy trả slice mới:

```go
func addDefault(values []int) []int {
    return append(values, 0)
}

values = addDefault(values)
```

## 3. Thao tác thường dùng

### Xóa phần tử tại chỉ số i, không giữ thứ tự

```go
values[i] = values[len(values)-1]
values = values[:len(values)-1]
```

### Xóa phần tử tại chỉ số i, giữ thứ tự

```go
copy(values[i:], values[i+1:])
values = values[:len(values)-1]
```

Phải kiểm tra `0 <= i && i < len(values)` trước khi xóa.

### Lọc tại chỗ

```go
result := values[:0]
for _, value := range values {
    if value%2 == 0 {
        result = append(result, value)
    }
}
values = result
```

## 4. Best practices

1. Dùng slice cho danh sách có độ dài thay đổi và tham số hàm xử lý danh sách.
2. Luôn gán lại kết quả `append`: `s = append(s, value)`.
3. Biết rõ khi nào các slice chia sẻ backing array; dùng `copy` khi cần dữ liệu độc lập.
4. Preallocate bằng `make([]T, 0, expectedSize)` khi biết gần đúng số phần tử sẽ thêm.
5. Không preallocate tùy tiện nếu chưa có bằng chứng cần tối ưu.
6. Dùng `len(s) == 0` thay vì phân biệt `nil` và empty slice, trừ khi API yêu cầu khác nhau.
7. Trả slice mới từ hàm nếu hàm có thể thay đổi độ dài.
8. Kiểm tra biên trước khi cắt hoặc truy cập theo chỉ số.
9. Với slice giữ dữ liệu lớn không còn cần, cân nhắc `clear(s)` hoặc gán phần tử về zero value trước khi thu ngắn.

## 5. Bài tập

### Mức dễ

#### Bài 1: Tổng và trung bình

Viết `sumAndAverage(values []int) (int, float64)`. Với slice rỗng, trả về `(0, 0)`.

#### Bài 2: Nhân đôi

Viết `double(values []int)` sửa trực tiếp từng phần tử. Kiểm tra với `[1, 2, 3]` thành `[2, 4, 6]`.

#### Bài 3: Tìm kiếm

Viết `indexOf(values []int, target int) int`, trả chỉ số đầu tiên hoặc `-1` nếu không tìm thấy.

#### Bài 4: Sao chép độc lập

Tạo bản sao của `[]string{"Go", "Java", "Rust"}`. Sửa bản sao và chứng minh slice gốc không đổi.

### Mức dễ đến vừa

#### Bài 5: Lọc số chẵn

Viết `filterEven(values []int) []int`. Kết quả không được làm thay đổi input.

Ví dụ: `[1, 2, 3, 4, 6]` thành `[2, 4, 6]`.

#### Bài 6: Xóa theo chỉ số

Viết `removeAt(values []int, index int) ([]int, bool)`:

- Giữ nguyên thứ tự.
- Trả `false` nếu chỉ số không hợp lệ.
- Không gây panic với slice rỗng.

#### Bài 7: Loại bỏ phần tử trùng

Viết `unique(values []int) []int`, giữ thứ tự xuất hiện đầu tiên.

Ví dụ: `[3, 1, 3, 2, 1]` thành `[3, 1, 2]`.

#### Bài 8: Chia thành nhóm

Viết `chunk(values []int, size int) [][]int`. Ví dụ với `[1,2,3,4,5]`, `size=2` trả `[[1,2], [3,4], [5]]`.

Quy định: nếu `size <= 0`, trả slice rỗng.

## 6. Gợi ý

<details>
<summary>Mở gợi ý</summary>

- Bài 1: xử lý `len(values) == 0` trước khi chia.
- Bài 4: dùng `copy` hoặc `append([]string(nil), source...)`.
- Bài 5: tạo kết quả với capacity bằng `len(values)`; chỉ append phần tử đạt điều kiện.
- Bài 6: dùng `copy(values[index:], values[index+1:])`, sau đó thu ngắn slice.
- Bài 7: dùng `map[int]bool` để ghi nhận phần tử đã gặp.
- Bài 8: tăng chỉ số theo `size`; giới hạn điểm cuối bằng `len(values)`.

</details>

## Checklist

- Tôi giải thích được `len` và `cap`.
- Tôi hiểu vì sao phải nhận lại kết quả `append`.
- Tôi biết khi nào hai slice chia sẻ backing array.
- Tôi sao chép được một slice độc lập.
- Tôi hoàn thành ít nhất 6/8 bài tập.

