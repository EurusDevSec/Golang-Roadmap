# 10 — Concurrency có kiểm soát

## Học trọng tâm

- Goroutine là đơn vị công việc đồng thời, không phải lý do để bỏ qua lifecycle.
- Channel dùng để truyền dữ liệu/ownership; mutex bảo vệ state chia sẻ.
- `sync.WaitGroup`, worker pool, semaphore và fan-out/fan-in.
- Race condition, deadlock, goroutine leak.

## Best practices

- Mọi goroutine phải có điều kiện kết thúc rõ ràng.
- Giới hạn concurrency; không tạo một goroutine cho mỗi item khi input không giới hạn.
- Owner đóng channel; receiver không đóng channel do người khác tạo.
- Không đóng channel chỉ để “dọn dẹp” nếu không cần báo kết thúc.
- Truyền context để hủy toàn bộ pipeline khi một bước thất bại.
- Chạy `go test -race ./...` cho code concurrent.
- Ưu tiên code tuần tự cho đến khi concurrency mang lại lợi ích đo được.

## Bài tập

1. Chạy 20 health check với tối đa 4 worker.
2. Thu kết quả về một channel và giữ đúng URL tương ứng.
3. Hủy toàn bộ batch khi context hết hạn.
4. Viết counter an toàn bằng mutex, xác nhận bằng race detector.
5. Tìm và sửa goroutine leak trong producer gửi vào channel không có receiver.

Checkpoint: concurrent checker không leak, có worker limit và kết quả ổn định.

