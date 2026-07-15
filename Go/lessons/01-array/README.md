# Array trong Go

## Mục tiêu

Sau bài này, bạn có thể:

- Khai báo, đọc, sửa và duyệt array.
- Hiểu array là kiểu có kích thước cố định.
- Biết khi nào dùng array và khi nào nên dùng slice.
- Tránh các lỗi phổ biến khi truyền array vào hàm.

## 1. Kiến thức cốt lõi

### Khai báo

```go
var scores [3]int                 // [0 0 0]
days := [3]string{"Mon", "Tue", "Wed"}
ids := [...]int{10, 20, 30}      // compiler tự suy ra độ dài là 3
```

Cú pháp kiểu array là `[N]T`, trong đó:

- `N`: số phần tử, phải xác định khi biên dịch.
- `T`: kiểu của mỗi phần tử.

`[3]int` và `[4]int` là hai kiểu khác nhau.

### Zero value

Nếu chưa gán giá trị, từng phần tử nhận zero value của kiểu:

```go
var numbers [3]int    // [0 0 0]
var flags [2]bool     // [false false]
var names [2]string   // ["" ""]
```

### Đọc và sửa phần tử

```go
numbers := [3]int{10, 20, 30}

fmt.Println(numbers[0]) // 10
numbers[1] = 99
fmt.Println(numbers)    // [10 99 30]
```

Chỉ số hợp lệ chạy từ `0` đến `len(numbers)-1`. Truy cập ngoài phạm vi sẽ gây lỗi.

### Duyệt array

Khi cần cả chỉ số và giá trị:

```go
for i, value := range numbers {
    fmt.Println(i, value)
}
```

Khi chỉ cần giá trị:

```go
for _, value := range numbers {
    fmt.Println(value)
}
```

Khi cần sửa trực tiếp từng phần tử, dùng chỉ số:

```go
for i := range numbers {
    numbers[i] *= 2
}
```

Biến `value` trong `for _, value := range numbers` là một bản sao. Gán lại `value` không sửa phần tử gốc.

## 2. Array được sao chép khi gán hoặc truyền vào hàm

```go
func change(a [3]int) {
    a[0] = 100
}

func main() {
    numbers := [3]int{1, 2, 3}
    change(numbers)
    fmt.Println(numbers) // [1 2 3]
}
```

Muốn sửa array gốc, truyền con trỏ:

```go
func change(a *[3]int) {
    a[0] = 100
}

change(&numbers)
```

Tuy nhiên, với dữ liệu có độ dài linh hoạt, thường nên nhận slice:

```go
func double(values []int) {
    for i := range values {
        values[i] *= 2
    }
}

numbers := [3]int{1, 2, 3}
double(numbers[:])
```

## 3. So sánh array

Hai array có cùng kiểu có thể so sánh bằng `==` nếu kiểu phần tử có thể so sánh:

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}

fmt.Println(a == b) // true
```

## 4. Khi nào nên dùng array?

Dùng array khi kích thước là một phần cố định của bài toán:

- Tọa độ RGB: `[3]uint8`.
- Ma trận cố định: `[3][3]int`.
- Mã định danh có đúng số byte: `[16]byte`.
- Bảng tra cứu có kích thước cố định.

Dùng slice `[]T` khi số lượng phần tử có thể thay đổi hoặc khi viết hàm dùng lại cho nhiều độ dài.

## 5. Best practices

1. Dùng `[...]T{...}` khi độ dài có thể suy ra từ dữ liệu khởi tạo.
2. Dùng `len(a)` thay cho viết cứng độ dài trong vòng lặp.
3. Dùng `range` để duyệt; dùng chỉ số khi cần sửa phần tử gốc.
4. Chỉ dùng array khi độ dài cố định mang ý nghĩa nghiệp vụ.
5. Hàm xử lý danh sách nói chung nên nhận slice `[]T`.
6. Nhớ rằng gán hoặc truyền array sẽ sao chép toàn bộ array.
7. Với array lớn cần chỉnh sửa tại chỗ, cân nhắc con trỏ `*[N]T` hoặc slice.
8. Đặt tên theo ý nghĩa dữ liệu, ví dụ `temperatures`, không đặt tên chung chung như `arr`.

## 6. Lỗi thường gặp

### Nhầm array với slice

```go
a := [3]int{1, 2, 3} // array, độ dài cố định
s := []int{1, 2, 3}  // slice, độ dài linh hoạt
```

Không thể dùng `append` trực tiếp với array:

```go
// a = append(a, 4) // sai
s = append(s, 4)    // đúng
```

### Sửa biến value trong range

```go
for _, value := range numbers {
    value *= 2 // không sửa numbers
}
```

Cách đúng:

```go
for i := range numbers {
    numbers[i] *= 2
}
```

### Viết cứng giới hạn vòng lặp

```go
for i := 0; i < 3; i++ { // khó bảo trì
    fmt.Println(numbers[i])
}
```

Nên viết:

```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

## 7. Bài tập

Quy ước: mỗi bài viết thành một chương trình Go độc lập. Làm bài trước khi xem gợi ý.

### Mức dễ

#### Bài 1: Khởi tạo và in array

Tạo array chứa 5 số nguyên `2, 4, 6, 8, 10`. In:

- Toàn bộ array.
- Phần tử đầu tiên và cuối cùng.
- Độ dài array.

Kết quả mong đợi:

```text
[2 4 6 8 10]
2 10
5
```

#### Bài 2: Tính tổng

Cho:

```go
numbers := [5]int{3, 7, 2, 9, 4}
```

Dùng `range` để tính và in tổng các phần tử. Kết quả: `25`.

#### Bài 3: Đếm số chẵn

Cho:

```go
numbers := [...]int{1, 2, 4, 7, 8, 11}
```

Đếm số phần tử chẵn. Kết quả: `3`.

#### Bài 4: Nhân đôi tại chỗ

Cho:

```go
numbers := [4]int{1, 3, 5, 7}
```

Sửa trực tiếp array để mỗi phần tử tăng gấp đôi. Kết quả: `[2 6 10 14]`.

Yêu cầu: dùng chỉ số, không tạo array thứ hai.

### Mức dễ đến vừa

#### Bài 5: Tìm lớn nhất và vị trí

Cho:

```go
numbers := [6]int{-3, 12, 5, 12, 1, 8}
```

In giá trị lớn nhất và chỉ số xuất hiện đầu tiên của nó.

Kết quả mong đợi:

```text
max=12 index=1
```

Không khởi tạo `max := 0`, vì cách đó sai khi tất cả phần tử đều âm.

#### Bài 6: Đảo ngược array

Cho:

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

Đảo ngược ngay trên array gốc. Kết quả: `[5 4 3 2 1]`.

Yêu cầu:

- Không tạo array phụ.
- Chỉ duyệt đến giữa array.

#### Bài 7: Hàm nhận array và tính trung bình

Viết hàm:

```go
func average(scores [5]float64) float64
```

Hàm trả về điểm trung bình. Kiểm tra với:

```go
scores := [5]float64{8, 7.5, 9, 6.5, 9}
```

Kết quả: `8`.

Sau đó trả lời: tại sao hàm này không nhận được `[4]float64`?

#### Bài 8: Xoay phải một vị trí

Cho:

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

Xoay các phần tử sang phải một vị trí. Phần tử cuối chuyển lên đầu.

Kết quả: `[5 1 2 3 4]`.

Yêu cầu:

- Sửa array tại chỗ.
- Chỉ dùng một biến tạm kiểu `int`.

## 8. Gợi ý

<details>
<summary>Mở gợi ý sau khi đã tự làm</summary>

- Bài 1: phần tử cuối có chỉ số `len(numbers)-1`.
- Bài 2: tạo `sum := 0`, rồi cộng từng `value`.
- Bài 3: số chẵn thỏa `value%2 == 0`.
- Bài 4: dùng `for i := range numbers` và sửa `numbers[i]`.
- Bài 5: khởi tạo `max := numbers[0]`; chỉ cập nhật khi gặp giá trị lớn hơn.
- Bài 6: đổi chỗ `numbers[i]` với `numbers[len(numbers)-1-i]`.
- Bài 7: `[5]float64` và `[4]float64` là hai kiểu khác nhau.
- Bài 8: lưu phần tử cuối, dịch từ phải sang trái, rồi gán biến tạm vào vị trí `0`.

</details>

## 9. Checklist tự đánh giá

Bạn đã nắm bài khi có thể trả lời “có” cho các câu sau:

- Tôi phân biệt được `[3]int` và `[]int`.
- Tôi biết zero value của một array.
- Tôi biết cách sửa phần tử khi dùng `range`.
- Tôi hiểu array được sao chép khi truyền vào hàm.
- Tôi biết khi nào nên chọn slice thay cho array.
- Tôi tự hoàn thành ít nhất 6/8 bài tập.
